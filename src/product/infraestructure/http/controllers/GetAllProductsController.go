package controllers

import (
	"api-order/src/product/application"
	"api-order/src/shared/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllProductsController struct {
	ProductService *application.GetAllProductsUseCase
}

func NewGetAllProductsController(productService *application.GetAllProductsUseCase) *GetAllProductsController {
	return &GetAllProductsController{ProductService: productService}
}

func (ctr *GetAllProductsController) Run(ctx *gin.Context){
	products, err := ctr.ProductService.Run()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "Error al obtener los productos",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Productos obtenidos",
		Data:    products,
		Error:   nil,
	})
}