package routes

import (
	motionhttp "api-order/src/data/motion/infrastructure/http" // Alias import
	"api-order/src/shared/middlewares"                         // Import middleware package

	"github.com/gin-gonic/gin"
)

// MotionRoutes configures routes for the motion feature
// Assumes placement under a /data prefix in the main router
func MotionRoutes(router *gin.RouterGroup) {
	// Initialize controller
	registerRecordController := motionhttp.SetUpRegisterRecordController()

	// Define the POST route for registering data.
	// Apply JWTAuthMiddleware for authentication.
	router.POST("/", middlewares.JWTAuthMiddleware(), registerRecordController.Run)
}
