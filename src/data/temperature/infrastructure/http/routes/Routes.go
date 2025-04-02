package routes

import (
	temperaturehttp "api-order/src/data/temperature/infrastructure/http" // Alias import
	"api-order/src/shared/middlewares"                                   // Import middleware package

	"github.com/gin-gonic/gin"
)

// TemperatureRoutes configures routes for the temperature feature
// It assumes these routes might live under a /data prefix in the main router
func TemperatureRoutes(router *gin.RouterGroup) {
	// Initialize controller
	registerRecordController := temperaturehttp.SetUpRegisterRecordController()

	// Define the POST route for registering data.
	// Apply JWTAuthMiddleware: Even though kit_id comes from the request,
	// this ensures only authenticated users can submit data.
	// You could add further validation later to check if the authenticated user owns the kit_id.
	router.POST("/", middlewares.JWTAuthMiddleware(), registerRecordController.Run)
}
