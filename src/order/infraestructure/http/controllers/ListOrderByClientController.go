package controllers

import (
	"api-order/src/order/application"
	"api-order/src/shared/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ListOrdersByClientController struct {
	OrderService *application.ListOrdersByClientUseCase
}

func NewListOrderByClientController(orderService *application.ListOrdersByClientUseCase) *ListOrdersByClientController {
	return &ListOrdersByClientController{OrderService: orderService}
}

func (ctr *ListOrdersByClientController) Run(ctx *gin.Context){
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10,64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Id inválido",
			Data: nil,
			Error: err.Error(),
		})
		return 
	}

	ordersClient, err := ctr.OrderService.Run(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message:"Error al obtener al cliente",
			Error: err.Error(),
			Data: nil,
		})

		return
	}


	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Cliente obtenido con éxito",
		Error: nil,
		Data: ordersClient,
	})

}