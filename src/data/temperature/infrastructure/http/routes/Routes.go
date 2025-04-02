package routes

import (
	temperaturehttp "api-order/src/data/temperature/infrastructure/http" // Alias import
	"api-order/src/shared/middlewares"                                   // Import middleware package

	"github.com/gin-gonic/gin"
)

func TemperatureRoutes(router *gin.RouterGroup) {
	// Initialize controllers
	registerRecordController := temperaturehttp.SetUpRegisterRecordController()
	getMinutesRecordsController := temperaturehttp.SetUpGetMinutesRecordsController() // <- Inicializar nuevo controller

	// POST / -> Register data (existing route)
	router.POST("/", middlewares.JWTAuthMiddleware(), registerRecordController.Run)

	// GET /kit/:kit_id/minutes/:minutes -> Get records for the last N minutes <- NUEVA RUTA
	router.GET("/kit/:kit_id/minutes/:minutes", middlewares.JWTAuthMiddleware(), getMinutesRecordsController.Run)
	// Note: Add authorization check if needed to ensure user owns the kit_id
}
