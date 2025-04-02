package controllers

import (
	"api-order/src/order/application"
	"api-order/src/order/infraestructure/http/request"
	"api-order/src/shared/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateOrderController struct {
	OrderService *application.CreateOrderUseCase
	Validator *validator.Validate
}

func NewCreateOrderController(orderService *application.CreateOrderUseCase) *CreateOrderController {
	return&CreateOrderController{OrderService: orderService, Validator: validator.New()}
}


func (ctr *CreateOrderController)Run(ctx *gin.Context){
	var req request.CreateOrderRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Llene todos los campos",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	order, err := ctr.OrderService.Run(req.Client_id, req.Product_id,req.Quantity, req.Total_price, req.Status, )

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error al crear la orden",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, responses.Response{
		Success: true,
		Message: "Orden creada",
		Data:    order,
		Error:   nil,
	})
}