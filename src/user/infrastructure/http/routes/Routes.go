package routes

import (
	"api-order/src/shared/middlewares" // Import authentication middleware
	"api-order/src/user/infrastructure/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	// Instantiate controllers using the setup functions
	registerController := http.SetUpRegisterUserController()
	loginController := http.SetUpLoginController()
	getUserController := http.SetUpGetUserByIdController()
	updateUserController := http.SetUpUpdateUserController()

	// Public routes
	router.POST("/", registerController.Run)   // Register User
	router.POST("/login", loginController.Run) // Login User

	// Protected routes (apply authentication middleware)
	// Create a subgroup for routes requiring authentication
	authorized := router.Group("/")
	authorized.Use(middlewares.JWTAuthMiddleware()) // Apply your JWT auth middleware
	{
		authorized.GET("/:id", getUserController.Run)    // Get User By ID
		authorized.PUT("/:id", updateUserController.Run) // Update User
	}
}
