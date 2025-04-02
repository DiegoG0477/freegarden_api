package controllers

import (
	"api-order/src/client/application"
	"api-order/src/client/infraestructure/http/controllers/helpers"
	"api-order/src/client/infraestructure/http/request"
	"api-order/src/shared/middlewares"
	"api-order/src/shared/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	ClientService *application.AuthUseCase
	BcrypttHelper *helpers.BcryptHelper
}

func NewAuthController(clientService *application.AuthUseCase) *AuthController {
	return &AuthController{ClientService: clientService}
}

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
		switch err.Error() {
		case "sql: no rows in result set":
			ctx.JSON(http.StatusNotFound, responses.Response{
				Success: false,
				Message: "El email no existe",
				Error:   err.Error(),
				Data:    nil,
			})
		default:
			ctx.JSON(http.StatusInternalServerError, responses.Response{
				Success: false,
				Message: "Error al iniciar sesión",
				Error:   err.Error(),
				Data:    nil,
			})
		}

		return
	}

	if err := ctr.BcrypttHelper.ComparePassword(client.Password, []byte(AuthRequest.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, responses.Response{
			Success: false,
			Message: "Contraseña incorrecta",
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}

	token, err := middlewares.GenerateJWT(int64(client.ID), client.Email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error al generar token",
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Sesión iniciada",
		Error:   nil,
		Data: map[string]interface{}{
			"token":    token,
			"Id":       client.ID,
			"LastName": client.LastName,
			"Name":     client.Name,
			"Email":    client.Email,
		},
	})
}
