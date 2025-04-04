package controllers

import (
	"api-order/src/shared/middlewares" // Assuming JWT generation is here
	"api-order/src/shared/responses"
	"api-order/src/user/application"
	"api-order/src/user/infrastructure/http/request"
	"database/sql"
	"errors"
	"net/http"
	"strings"

	// Use alias if needed to avoid conflict if client also has entities
	userEntities "api-order/src/user/domain/entities"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt" // Needed for bcrypt.ErrMismatchedHashAndPassword
)

// LoginResponseData defines the structure for the successful login response data
type LoginResponseData struct {
	Token string                    `json:"token"`
	User  userEntities.UserResponse `json:"user"`
}

type LoginController struct {
	UserService *application.LoginUseCase
	Validator   *validator.Validate
	// BcryptHelper is used within the use case now
}

func NewLoginController(userService *application.LoginUseCase) *LoginController {
	return &LoginController{
		UserService: userService,
		Validator:   validator.New(),
	}
}

// @Summary      Authenticate a user
// @Description  Logs in a user using email and password, returns user details and a JWT token upon success.
// @Tags         Users Authentication
// @Accept       json
// @Produce      json
// @Param        credentials body request.LoginRequest true "User Login Credentials"
// @Success      200  {object}  responses.Response{data=LoginResponseData} "Login successful"
// @Failure      400  {object}  responses.Response "Invalid request body format or validation failed"
// @Failure      401  {object}  responses.Response "Incorrect password or invalid credentials"
// @Failure      404  {object}  responses.Response "Email not found"
// @Failure      500  {object}  responses.Response "Internal server error during login or token generation"
// @Router       /v1/users/login [post]
func (ctr *LoginController) Run(ctx *gin.Context) {
	var req request.LoginRequest

	// Bind and validate JSON input
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false, Message: "Error procesando la solicitud. Verifique los campos.", Error: err.Error(), Data: nil,
		})
		return
	}
	if err := ctr.Validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false, Message: "Datos de login inválidos.", Error: err.Error(), Data: nil,
		})
		return
	}

	// Execute login use case
	user, err := ctr.UserService.Run(req.Email, req.Password)

	// Handle errors from use case
	if err != nil {
		// Check for "not found" error (specific check for sql.ErrNoRows is good)
		if errors.Is(err, sql.ErrNoRows) || strings.Contains(err.Error(), "not found") { // Check message if Is doesn't work across layers
			ctx.JSON(http.StatusNotFound, responses.Response{
				Success: false, Message: "El email no está registrado.", Error: "Email not found", Data: nil,
			})
			return
		}
		// Check for invalid credentials / password mismatch
		// Compare directly with bcrypt error or the specific error string from the use case
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) || (err != nil && err.Error() == "invalid credentials") {
			ctx.JSON(http.StatusUnauthorized, responses.Response{
				Success: false, Message: "Contraseña incorrecta.", Error: "Invalid credentials", Data: nil,
			})
			return
		}

		// Handle other internal errors
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false, Message: "Error al intentar iniciar sesión.", Error: "Internal server error", Data: nil,
		})
		return
	}

	// Generate JWT Token
	// Assuming GenerateJWT takes user ID (int64) and email (string)
	token, err := middlewares.GenerateJWT(user.ID, user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false, Message: "Error al generar el token de sesión.", Error: "Failed to generate token", Data: nil,
		})
		return
	}

	// Prepare successful response data
	responseData := LoginResponseData{
		Token: token,
		User:  user.ToResponse(), // Use the response struct without password
	}

	// Send success response
	ctx.JSON(http.StatusOK, responses.Response{
		Success: true, Message: "Sesión iniciada con éxito.", Error: nil, Data: responseData,
	})
}
