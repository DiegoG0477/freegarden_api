package http

import (
	"api-order/src/data/airquality/application"                     // Adjusted import path
	"api-order/src/data/airquality/domain/ports"                    // Adjusted import path
	"api-order/src/data/airquality/infrastructure/adapters"         // Adjusted import path
	"api-order/src/data/airquality/infrastructure/http/controllers" // Adjusted import path
	"log"
)

// Declare repository variable specific to air quality data
var (
	airQualityRepository ports.IAirQualityData
)

// Initialize air quality dependencies
func InitializeAirQualityDependencies() {
	var err error
	airQualityRepository, err = adapters.NewAirQualityDataRepositoryMysql()
	if err != nil {
		log.Fatalf("Error initializing air quality data repository: %v", err)
	}
}

// Setup function for RegisterRecordController
func SetUpRegisterRecordController() *controllers.RegisterRecordController {
	// Ensure dependencies are initialized
	if airQualityRepository == nil {
		InitializeAirQualityDependencies()
	}
	registerRecordService := application.NewRegisterRecordUseCase(airQualityRepository)
	return controllers.NewRegisterRecordController(registerRecordService)
}
