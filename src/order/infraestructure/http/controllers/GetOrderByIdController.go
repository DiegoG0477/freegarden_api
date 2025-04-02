package controllers

import (
	"api-order/src/order/application"
	"api-order/src/shared/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetOrderByIdController struct {
	OrderService *application.GetOrderByIdUseCase
}

func NewGetOrderByIdController(orderService *application.GetOrderByIdUseCase) *GetOrderByIdController {
	return &GetOrderByIdController{OrderService: orderService}
}

func (ctr *GetOrderByIdController) Run(ctx *gin.Context){
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest,responses.Response{
			Success: false,
			Message: "Id inválido",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	order, err := ctr.OrderService.Run(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error al obtener la orden",
			Error: err.Error(),
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusAccepted, responses.Response{
		Success: true,
		Message: "Orden obtenida con éxito",
		Error: nil,
		Data: order,
	})
}