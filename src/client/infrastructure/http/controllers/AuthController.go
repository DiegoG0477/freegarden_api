package controllers

import (
	"api-order/src/client/application"
	"api-order/src/client/infrastructure/http/controllers/helpers"
	"api-order/src/client/infrastructure/http/request"
	"api-order/src/shared/middlewares"
	"api-order/src/shared/responses"
	"net/http"

	_ "api-order/src/client/domain/entities" // Necesario para documentar la estructura interna de data en Success

	"github.com/gin-gonic/gin"
)

// AuthResponseData defines the structure for the successful login response data
// Necesitamos definir esto porque map[string]interface{} no es muy descriptivo en Swagger.
type AuthResponseData struct {
	Token    string `json:"token"`
	Id       int    `json:"Id"`
	LastName string `json:"LastName"`
	Name     string `json:"Name"`
	Email    string `json:"Email"`
}

type AuthController struct {
	ClientService *application.AuthUseCase
	BcrypttHelper *helpers.BcryptHelper
}

func NewAuthController(clientService *application.AuthUseCase) *AuthController {
	// Inicializa BcrypttHelper aquí si aún no lo está. Asumiendo que se inicializa externamente por ahora.
	// Si no, añade: bcryptHelper, _ := helpers.NewBcryptHelper()
	return &AuthController{
		ClientService: clientService,
		BcrypttHelper: &helpers.BcryptHelper{}, // Asumiendo que BcryptHelper es stateless o inicializado
	}
}

// @Summary      Authenticate a client
// @Description  Logs in a client using email and password, returns a JWT token upon success.
// @Tags         Clients Authentication
// @Accept       json
// @Produce      json
// @Param        credentials body request.AuthRequest true "Client Login Credentials"
// @Success      200  {object}  responses.Response{data=AuthResponseData} "Login successful"
// @Failure      400  {object}  responses.Response "Invalid request body format"
// @Failure      401  {object}  responses.Response "Incorrect password"
// @Failure      404  {object}  responses.Response "Email not found"
// @Failure      500  {object}  responses.Response "Internal server error during login or token generation"
// @Router       /clients/auth [post]
func (ctr *AuthController) Run(ctx *gin.Context) {
	var AuthRequest request.AuthRequest

	if err := ctx.BindJSON(&AuthRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Error al procesar la solicitud",
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}

	client, err := ctr.ClientService.Run(AuthRequest.Email)

	if err != nil {
		// Usar sql.ErrNoRows para una comparación más robusta si es posible
		if err.Error() == "sql: no rows in result set" { // O if errors.Is(err, sql.ErrNoRows)
			ctx.JSON(http.StatusNotFound, responses.Response{
				Success: false,
				Message: "El email no existe",
				Error:   "Email not found", // Mensaje de error más estándar para API
				Data:    nil,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, responses.Response{
				Success: false,
				Message: "Error al iniciar sesión",
				Error:   "Internal server error", // No exponer detalles internos
				Data:    nil,
			})
		}
		return
	}

	// Asegúrate que BcrypttHelper esté inicializado
	if ctr.BcrypttHelper == nil {
		// Manejar el caso donde BcrypttHelper no está listo, quizás loguear y retornar 500
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false, Message: "Error interno del servidor", Error: "Bcrypt helper not initialized", Data: nil,
		})
		return
	}

	if err := ctr.BcrypttHelper.ComparePassword(client.Password, []byte(AuthRequest.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, responses.Response{
			Success: false,
			Message: "Contraseña incorrecta",
			Error:   "Invalid credentials", // Mensaje estándar
			Data:    nil,
		})
		return
	}

	token, err := middlewares.GenerateJWT(int64(client.ID), client.Email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error al generar token",
			Error:   "Failed to generate token",
			Data:    nil,
		})
		return
	}

	// Usar la struct definida para la respuesta
	responseData := AuthResponseData{
		Token:    token,
		Id:       client.ID,
		LastName: client.LastName,
		Name:     client.Name,
		Email:    client.Email,
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Sesión iniciada",
		Error:   nil,
		Data:    responseData,
	})
}
