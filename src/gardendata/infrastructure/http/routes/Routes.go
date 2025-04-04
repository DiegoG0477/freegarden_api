package routes

import (
	"api-order/src/gardendata/infrastructure/http" // Corrected path
	"api-order/src/shared/middlewares"             // Assuming common auth middleware

	"github.com/gin-gonic/gin"
)

func GardenDataRoutes(router *gin.RouterGroup) {
	// Instantiate controllers using the setup functions
	registerController := http.SetUpRegisterGardenDataController()
	getController := http.SetUpGetMinutesGardenDataController()

	// Apply authentication middleware if needed for these routes
	// Data ingestion (POST) might use API keys, GET might use user tokens
	// Applying standard user auth middleware here for consistency example:

	// Define routes
	router.POST("/", registerController.Run)                                                        // Register new data
	router.GET("/kit/:kit_id/minutes/:minutes", middlewares.JWTAuthMiddleware(), getController.Run) // Get recent data
}
