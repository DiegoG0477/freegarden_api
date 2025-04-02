package controllers

import (
	"api-order/src/client/application"
	"api-order/src/client/infrastructure/http/request"
	"api-order/src/shared/responses"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateClientController struct {
	ClientService *application.CreateClientUseCase
	Validator     *validator.Validate
}

func NewCreateClientController(clientService *application.CreateClientUseCase) *CreateClientController {
	return &CreateClientController{
		ClientService: clientService,
		Validator:     validator.New(),
	}
}

// @Summary      Create a new client
// @Description  Registers a new client in the system.
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        client body request.CreateClientRequest true "Client data to create"
// @Success      201  {object}  responses.Response{data=entities.Client} "Client created successfully"
// @Failure      400  {object}  responses.Response "Invalid request body or validation failed"
// @Failure      409  {object}  responses.Response "Email already exists"
// @Failure      500  {object}  responses.Response "Internal server error"
// @Router       /clients/ [post]
func (ctr *CreateClientController) Run(ctx *gin.Context) {
	var req request.CreateClientRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Llene todos los campos",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	if err := ctr.Validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Los datos enviados no son válidos",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	client, err := ctr.ClientService.Run(req.Name, req.Email, req.Password, req.LastName)

	if err != nil {
		if strings.Contains(err.Error(), "unique_client_email") {
			ctx.JSON(http.StatusConflict, responses.Response{
				Success: false,
				Message: "El email ya existe",
				Data:    nil,
				Error:   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Error al crear el cliente",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, responses.Response{
		Success: true,
		Message: "Cliente creado correctamente",
		Data:    client,
		Error:   nil,
	})
}
