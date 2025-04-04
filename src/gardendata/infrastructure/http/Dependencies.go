package http

import (
	// Standard library imports if needed (e.g., "log")
	"log"

	"api-order/src/gardendata/application" // Corrected paths
	"api-order/src/gardendata/domain/ports"
	"api-order/src/gardendata/infrastructure/adapters"
	"api-order/src/gardendata/infrastructure/http/controllers"
)

// Variables holding instances (consider central DI instead of package vars)
var (
	gardenDataRepository ports.IGardenData
	// Use cases
	registerGardenDataUseCase   *application.RegisterGardenDataUseCase
	getMinutesGardenDataUseCase *application.GetMinutesGardenDataUseCase
)

// Initialize dependencies for the GardenData feature
// WARNING: init() functions in multiple packages can have unpredictable order
// and make dependencies harder to manage. Prefer explicit initialization.
func init() {
	var err error
	gardenDataRepository, err = adapters.NewGardenDataRepositoryMysql()
	if err != nil {
		// Use log.Fatalf to stop execution if critical dependencies fail
		log.Fatalf("Error initializing GardenData repository: %v", err)
	}

	// Initialize Use Cases
	registerGardenDataUseCase = application.NewRegisterGardenDataUseCase(gardenDataRepository)
	getMinutesGardenDataUseCase = application.NewGetMinutesGardenDataUseCase(gardenDataRepository)
}

// Setup functions for GardenData controllers

func SetUpRegisterGardenDataController() *controllers.RegisterGardenDataController {
	// Ensure use case is initialized
	if registerGardenDataUseCase == nil {
		log.Fatal("RegisterGardenDataUseCase not initialized") // Or handle error appropriately
	}
	return controllers.NewRegisterGardenDataController(registerGardenDataUseCase)
}

func SetUpGetMinutesGardenDataController() *controllers.GetMinutesGardenDataController {
	// Ensure use case is initialized
	if getMinutesGardenDataUseCase == nil {
		log.Fatal("GetMinutesGardenDataUseCase not initialized") // Or handle error appropriately
	}
	return controllers.NewGetMinutesGardenDataController(getMinutesGardenDataUseCase)
}
