package routes

import (
	airqualityhttp "api-order/src/data/airquality/infrastructure/http" // Alias import
	"api-order/src/shared/middlewares"

	"github.com/gin-gonic/gin"
)

func AirQualityRoutes(router *gin.RouterGroup) {
	// Initialize controllers
	registerRecordController := airqualityhttp.SetUpRegisterRecordController()
	getMinutesRecordsController := airqualityhttp.SetUpGetMinutesRecordsController() // <- Inicializar nuevo controller

	// POST / -> Register data (existing route)
	router.POST("/", middlewares.JWTAuthMiddleware(), registerRecordController.Run)

	// GET /kit/:kit_id/minutes/:minutes -> Get records for the last N minutes <- NUEVA RUTA
	router.GET("/kit/:kit_id/minutes/:minutes", middlewares.JWTAuthMiddleware(), getMinutesRecordsController.Run)
}
