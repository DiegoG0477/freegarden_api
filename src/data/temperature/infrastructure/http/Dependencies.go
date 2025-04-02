package http

import (
	"api-order/src/data/temperature/application"
	"api-order/src/data/temperature/domain/ports"
	"api-order/src/data/temperature/infrastructure/adapters"
	"api-order/src/data/temperature/infrastructure/http/controllers"
	"log"
)

// Declare repository variable specific to temperature data
var (
	temperatureRepository ports.ITemperatureData
)

// Initialize temperature dependencies
func InitializeTemperatureDependencies() {
	var err error
	temperatureRepository, err = adapters.NewTemperatureDataRepositoryMysql()
	if err != nil {
		log.Fatalf("Error initializing temperature data repository: %v", err)
	}
}

// Setup function for RegisterRecordController
func SetUpRegisterRecordController() *controllers.RegisterRecordController {
	// Ensure dependencies are initialized
	if temperatureRepository == nil {
		InitializeTemperatureDependencies()
	}
	registerRecordService := application.NewRegisterRecordUseCase(temperatureRepository)
	return controllers.NewRegisterRecordController(registerRecordService)
}
