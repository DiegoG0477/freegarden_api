package controllers

import (
	"api-order/src/order/application"
	"api-order/src/order/infraestructure/http/request"
	"api-order/src/shared/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateStatusController struct {
	OrderService *application.UpdateStatusOrderUseCase
}

func NewUpdateStatusController(orderService *application.UpdateStatusOrderUseCase) *UpdateStatusController{
	return &UpdateStatusController{OrderService: orderService}
}

func (ctr *UpdateStatusController) Run(ctx *gin.Context){
	var req request.UpdateOrderRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Llene todos los campos",
			Data: nil,
			Error: err.Error(),
		})
		return
	}

	updated, err := ctr.OrderService.Run(int64(req.ID), req.Status)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error al actualizar la orden",
			Data: nil,
			Error: err.Error(),
		})
	}


	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Orden actualizada con éxito",
		Error: nil,
		Data: updated,
	})
}