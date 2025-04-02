package controllers

import (
	"api-order/src/order/application"
	"api-order/src/shared/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteOrderController struct {
	OrderService *application.DeleteOrderUseCase
}

func NewDeleteOrderController(orderService *application.DeleteOrderUseCase) *DeleteOrderController{
	return &DeleteOrderController{OrderService: orderService}
}

func (ctr *DeleteOrderController) Run(ctx *gin.Context){
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Id inválido",
			Data: nil,
			Error: err.Error(),
		})
		return
	}


	deleted, err := ctr.OrderService.Run(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error al eliminar la orden",
			Error: err.Error(),
			Data: nil,
		})
		return
	}


	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Orden eliminada con éxito",
		Error: nil,
		Data: deleted,
	})

}

