package controllers

import (
	"api-order/src/product/application"
	"api-order/src/product/infraestructure/http/request"
	"api-order/src/shared/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	ProductService *application.CreateProductUseCase
}

func NewCreateProductController(productService *application.CreateProductUseCase) *CreateProductController {
	return &CreateProductController{ProductService: productService}
}

func (ctr *CreateProductController) CreateProduct(ctx *gin.Context) {
	var req request.ProductRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Llene todos los campos",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	product, err := ctr.ProductService.Run(req.Name, req.Description, req.Price, req.Stock)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error al crear el producto",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, responses.Response{
		Success: true,
		Message: "Producto creado",
		Data:    product,
		Error:   nil,
	})
}