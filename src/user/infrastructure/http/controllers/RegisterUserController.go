package controllers

import (
	"api-order/src/shared/responses"
	"api-order/src/user/application"
	"api-order/src/user/infrastructure/http/request"
	"errors"
	"net/http"

	// "strings" // Needed if checking generic duplicate errors

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegisterUserController struct {
	UserService *application.RegisterUserUseCase
	Validator   *validator.Validate
}

func NewRegisterUserController(userService *application.RegisterUserUseCase) *RegisterUserController {
	return &RegisterUserController{
		UserService: userService,
		Validator:   validator.New(),
	}
}

// @Summary      Register a new user
// @Description  Registers a new user with validation against an existing kit code. The kit code must NOT exist in the kits table.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user body request.RegisterUserRequest true "User Registration Data"
// @Success      201  {object}  responses.Response{data=entities.UserResponse} "User registered successfully"
// @Failure      400  {object}  responses.Response "Invalid request body or validation failed"
// @Failure      409  {object}  responses.Response "Conflict - Email already exists OR Kit Code already exists"
// @Failure      500  {object}  responses.Response "Internal server error during registration"
// @Router       /v1/users/ [post]
func (ctr *RegisterUserController) Run(ctx *gin.Context) {
	var req request.RegisterUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Error procesando la solicitud. Verifique los campos.",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	// Validate struct fields
	if err := ctr.Validator.Struct(req); err != nil {
		// Customize validation error messages if needed
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Datos inválidos proporcionados.",
			Data:    nil,
			Error:   err.Error(), // Consider formatting validation errors nicely
		})
		return
	}

	// Run the use case
	user, err := ctr.UserService.Run(req.FirstName, req.LastName, req.Email, req.Password, req.KitCode)

	if err != nil {
		// Handle specific domain errors
		if errors.Is(err, application.ErrKitCodeExists) {
			ctx.JSON(http.StatusConflict, responses.Response{
				Success: false,
				Message: "El código de kit proporcionado ya está en uso.",
				Data:    nil,
				Error:   err.Error(),
			})
			return
		}
		if errors.Is(err, application.ErrUserEmailExists) {
			ctx.JSON(http.StatusConflict, responses.Response{
				Success: false,
				Message: "El email proporcionado ya está registrado.",
				Data:    nil,
				Error:   err.Error(),
			})
			return
		}
		// Handle generic duplicate errors from DB if not caught by specific checks
		// Example: depends on the exact error string from your DB driver
		// if strings.Contains(err.Error(), "Duplicate entry") || strings.Contains(err.Error(), "UNIQUE constraint failed") {
		// 	ctx.JSON(http.StatusConflict, responses.Response{
		// 		Success: false,
		// 		Message: "El email ya existe.",
		// 		Data:    nil,
		// 		Error:   "Email already registered",
		// 	})
		// 	return
		// }

		// Handle other potential errors
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error al registrar el usuario.",
			Data:    nil,
			Error:   "Internal server error", // Don't expose internal details err.Error(),
		})
		return
	}

	// Return success response (without password)
	ctx.JSON(http.StatusCreated, responses.Response{
		Success: true,
		Message: "Usuario registrado correctamente.",
		Data:    user.ToResponse(), // Use the dedicated response struct
		Error:   nil,
	})
}
