package http

import (
	"api-order/src/kit/application"
	"api-order/src/kit/domain/ports"
	"api-order/src/kit/infrastructure/adapters"
	"api-order/src/kit/infrastructure/http/controllers"
	"log"
)

// Declare repository variable specific to kit
var (
	kitRepository ports.IKit
)

// Initialize kit dependencies. You might merge this with the client's init
// or keep them separate if preferred. Let's keep it separate for clarity.
func InitializeKitDependencies() {
	var err error
	kitRepository, err = adapters.NewKitRepositoryMysql()
	if err != nil {
		log.Fatalf("Error initializing kit repository: %v", err)
	}
	// Add other kit-specific dependencies here if needed in the future
}

// Setup function for CreateKitController
func SetUpCreateKitController() *controllers.CreateKitController {
	// Ensure dependencies are initialized before setting up controllers
	if kitRepository == nil {
		InitializeKitDependencies()
	}
	createKitService := application.NewCreateKitUseCase(kitRepository)
	return controllers.NewCreateKitController(createKitService)
}

// Setup function for GetKitsController
func SetUpGetKitsController() *controllers.GetKitsController {
	// Ensure dependencies are initialized
	if kitRepository == nil {
		InitializeKitDependencies()
	}
	getKitsService := application.NewGetKitsUseCase(kitRepository)
	return controllers.NewGetKitsController(getKitsService)
}
