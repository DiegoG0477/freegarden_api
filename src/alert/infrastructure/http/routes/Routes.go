package routes

import (
	alerthttp "api-order/src/alert/infrastructure/http" // Alias import
	"api-order/src/shared/middlewares"                  // Import middleware package

	"github.com/gin-gonic/gin"
)

// AlertRoutes configures routes for the alerts feature
func AlertRoutes(router *gin.RouterGroup) {
	// Initialize controllers
	registerAlertController := alerthttp.SetUpRegisterAlertController()
	getAlertsController := alerthttp.SetUpGetAlertsByKitIDController()

	// POST / -> Register a new alert
	router.POST("/", registerAlertController.Run)
	// GET /kit/:kit_id -> Get alerts for a specific kit
	router.GET("/:kit_id", middlewares.JWTAuthMiddleware(), getAlertsController.Run)
	// Note: Further authorization could be added here or in the use case/controller
	// to verify if the authenticated user owns the requested kit_id.
}
