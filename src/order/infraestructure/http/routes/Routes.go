package routes

import (
	"api-order/src/order/infraestructure/http"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.RouterGroup){
	createOrderController := http.SetUpCreate()
	listOrdersByClientController := http.ListOrdersByClientController()
	getOrderByIdController := http.GetByIdController()
	deleteOrderController := http.Delete()
	updateStatusController := http.UpdateStatusController()

	router.POST("/", createOrderController.Run)
	router.GET("/client/:id", listOrdersByClientController.Run)
	router.GET("/:id", getOrderByIdController.Run)
	router.DELETE("/:id", deleteOrderController.Run)
	router.PUT("/", updateStatusController.Run)

}