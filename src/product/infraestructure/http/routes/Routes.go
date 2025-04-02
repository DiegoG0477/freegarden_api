package routes

import (
	"api-order/src/product/infraestructure/http"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.RouterGroup){
	createProductController := http.SetUpCreateProduct()
	getAllProductsController := http.SetUpGetAllProducts()

	router.POST("/", createProductController.CreateProduct)
	router.GET("/", getAllProductsController.Run)
}