package http

import (
	kit "api-order/src/kit/domain/ports"
	kitAdpt "api-order/src/kit/infrastructure/adapters"
	"api-order/src/user/application"
	"api-order/src/user/application/services"
	"api-order/src/user/domain/ports"
	"api-order/src/user/infrastructure/adapters"
	"api-order/src/user/infrastructure/http/controllers"
	"api-order/src/user/infrastructure/http/controllers/helpers" // User's helpers
	"log"
)

var (
	userRepository ports.IUser
	kitRepository  kit.IKit
	encryptService services.IEncrypt // User's encrypt service interface
)

// Initialize dependencies for the User feature
// Note: This init might run alongside the client's init if both packages are imported.
// Consider a central dependency management approach.
func init() {
	var err error
	userRepository, err = adapters.NewUserRepositoryMysql()
	if err != nil {
		log.Fatalf("Error initializing user repository: %v", err)
	}

	kitRepository, err = kitAdpt.NewKitRepositoryMysql()
	if err != nil {
		log.Fatalf("Error initializing kit repository: %v", err)
	}

	// Use the user's helper for IEncrypt interface
	encryptService, err = helpers.NewBcryptHelper()
	if err != nil {
		log.Fatalf("Error initializing user encrypt service: %v", err)
	}
}

// Setup functions for User controllers

func SetUpRegisterUserController() *controllers.RegisterUserController {
	registerUseCase := application.NewRegisterUserUseCase(userRepository, kitRepository, encryptService)
	return controllers.NewRegisterUserController(registerUseCase)
}

func SetUpLoginController() *controllers.LoginController {
	loginUseCase := application.NewLoginUseCase(userRepository, encryptService)
	return controllers.NewLoginController(loginUseCase)
}

func SetUpGetUserByIdController() *controllers.GetUserByIdController {
	getUserUseCase := application.NewGetUserByIdUseCase(userRepository)
	return controllers.NewGetUserByIdController(getUserUseCase)
}

func SetUpUpdateUserController() *controllers.UpdateUserController {
	updateUserUseCase := application.NewUpdateUserUseCase(userRepository)
	return controllers.NewUpdateUserController(updateUserUseCase)
}
