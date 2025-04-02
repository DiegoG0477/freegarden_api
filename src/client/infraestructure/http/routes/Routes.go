package routes

import (
	"api-order/src/client/infraestructure/http"

	"github.com/gin-gonic/gin"
)

func ClientRoutes(router *gin.RouterGroup){
	createClientController := http.SetUpCreateClientController()
	authController := http.AuthControllers()
	getClientByIdController := http.GetClientByIdController()

	router.POST("/", createClientController.Run)
	router.POST("/auth", authController.Run)
	router.GET("/:id", getClientByIdController.Run)
}