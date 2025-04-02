package routes

import (
	airqualityhttp "api-order/src/data/airquality/infrastructure/http" // Alias import
	"api-order/src/shared/middlewares"                                 // Import middleware package

	"github.com/gin-gonic/gin"
)

// AirQualityRoutes configures routes for the air quality feature
// Assumes placement under a /data prefix in the main router
func AirQualityRoutes(router *gin.RouterGroup) {
	// Initialize controller
	registerRecordController := airqualityhttp.SetUpRegisterRecordController()

	// Define the POST route for registering data.
	// Apply JWTAuthMiddleware for authentication.
	router.POST("/", middlewares.JWTAuthMiddleware(), registerRecordController.Run)
}
