package routes

import (
	lighthttp "api-order/src/data/light/infrastructure/http" // Alias import
	"api-order/src/shared/middlewares"                       // Import middleware package

	"github.com/gin-gonic/gin"
)

// LightRoutes configures routes for the light feature
// Assumes placement under a /data prefix in the main router
func LightRoutes(router *gin.RouterGroup) {
	// Initialize controller
	registerRecordController := lighthttp.SetUpRegisterRecordController()

	// Define the POST route for registering data.
	// Apply JWTAuthMiddleware for authentication.
	router.POST("/", middlewares.JWTAuthMiddleware(), registerRecordController.Run)
}
