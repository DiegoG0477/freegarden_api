package routes

import (
	kithttp "api-order/src/kit/infrastructure/http" // Alias import
	"api-order/src/shared/middlewares"              // Import middleware package

	"github.com/gin-gonic/gin"
)

// KitRoutes configures routes for the kit feature
func KitRoutes(router *gin.RouterGroup) {
	// Initialize controllers using the setup functions from Dependencies.go
	createKitController := kithttp.SetUpCreateKitController()
	getKitsController := kithttp.SetUpGetKitsController()

	// Apply JWTAuthMiddleware to protect these routes
	// The middleware runs first, setting 'datUser' in context if valid
	router.POST("/", middlewares.JWTAuthMiddleware(), createKitController.Run)
	router.GET("/", middlewares.JWTAuthMiddleware(), getKitsController.Run)
}
