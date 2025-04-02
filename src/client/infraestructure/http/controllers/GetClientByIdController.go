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

func (ctr *GetClientByIdController) Run(ctx *gin.Context){
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10,64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest,responses.Response{
			Success: false,
			Message: "Id inválido",
			Data: nil,
			Error: err.Error(),
		})
        return
	}

	client, err := ctr.ClientService.Run(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error al obtener el cliente",
			Error: err.Error(),
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Cliente obtenido con éxito",
		Error: nil,
		Data: client,
	})
}