package http

import (
	"api-order/src/data/light/application"                     // Adjusted import path
	"api-order/src/data/light/domain/ports"                    // Adjusted import path
	"api-order/src/data/light/infrastructure/adapters"         // Adjusted import path
	"api-order/src/data/light/infrastructure/http/controllers" // Adjusted import path
	"log"
)

// Declare repository variable specific to light data
var (
	lightRepository ports.ILightData
)

// Initialize light dependencies
func InitializeLightDependencies() {
	var err error
	lightRepository, err = adapters.NewLightDataRepositoryMysql()
	if err != nil {
		log.Fatalf("Error initializing light data repository: %v", err)
	}
}

// Setup function for RegisterRecordController
func SetUpRegisterRecordController() *controllers.RegisterRecordController {
	// Ensure dependencies are initialized
	if lightRepository == nil {
		InitializeLightDependencies()
	}
	registerRecordService := application.NewRegisterRecordUseCase(lightRepository)
	return controllers.NewRegisterRecordController(registerRecordService)
}
