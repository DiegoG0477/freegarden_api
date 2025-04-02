package routes

import (
	motionhttp "api-order/src/data/motion/infrastructure/http" // Alias import
	"api-order/src/shared/middlewares"

	"github.com/gin-gonic/gin"
)

func MotionRoutes(router *gin.RouterGroup) {
	// Initialize controllers
	registerRecordController := motionhttp.SetUpRegisterRecordController()
	getMinutesRecordsController := motionhttp.SetUpGetMinutesRecordsController() // <- Inicializar nuevo controller

	// POST / -> Register data (existing route)
	router.POST("/", middlewares.JWTAuthMiddleware(), registerRecordController.Run)

	// GET /kit/:kit_id/minutes/:minutes -> Get records for the last N minutes <- NUEVA RUTA
	router.GET("/kit/:kit_id/minutes/:minutes", middlewares.JWTAuthMiddleware(), getMinutesRecordsController.Run)
}
