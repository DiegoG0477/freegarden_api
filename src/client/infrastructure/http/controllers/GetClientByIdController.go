package controllers

import (
	"api-order/src/client/application"
	"api-order/src/shared/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetClientByIdController struct {
	ClientService *application.GetClientByIdUseCase
}

func NewGetClientByIdController(clientService *application.GetClientByIdUseCase) *GetClientByIdController {
	return &GetClientByIdController{ClientService: clientService}
}

// @Summary      Get client by ID
// @Description  Retrieves the details of a specific client by their ID. This endpoint might be public or require specific permissions not covered by standard user JWT.
// @Tags         Clients
// @Produce      json
// @Param        id path int true "Client ID" Format(int64)
// @Success      200  {object}  responses.Response{data=entities.Client} "Client retrieved successfully"
// @Failure      400  {object}  responses.Response "Invalid client ID provided"
// @Failure      404  {object}  responses.Response "Client not found"
// @Failure      500  {object}  responses.Response "Internal server error while retrieving client"
// @Router       /clients/{id} [get]
func (ctr *GetClientByIdController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Id inválido",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	client, err := ctr.ClientService.Run(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error al obtener el cliente",
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Cliente obtenido con éxito",
		Error:   nil,
		Data:    client,
	})
}
