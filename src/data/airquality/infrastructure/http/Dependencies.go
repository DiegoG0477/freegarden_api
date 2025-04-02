package http

import (
	"api-order/src/data/airquality/application"
	"api-order/src/data/airquality/domain/ports"
	"api-order/src/data/airquality/infrastructure/adapters"
	"api-order/src/data/airquality/infrastructure/http/controllers"
	"log"
)

var (
	airQualityRepository ports.IAirQualityData
)

func InitializeAirQualityDependencies() {
	var err error
	if airQualityRepository == nil {
		airQualityRepository, err = adapters.NewAirQualityDataRepositoryMysql()
		if err != nil {
			log.Fatalf("Error initializing air quality data repository: %v", err)
		}
	}
}

// Setup for RegisterRecordController (unchanged)
func SetUpRegisterRecordController() *controllers.RegisterRecordController {
	if airQualityRepository == nil {
		InitializeAirQualityDependencies()
	}
	registerRecordService := application.NewRegisterRecordUseCase(airQualityRepository)
	return controllers.NewRegisterRecordController(registerRecordService)
}

// Setup function for GetMinutesRecordsController <- NUEVO SETUP
func SetUpGetMinutesRecordsController() *controllers.GetMinutesRecordsController {
	if airQualityRepository == nil {
		InitializeAirQualityDependencies()
	}
	getMinutesRecordService := application.NewGetMinutesRecordsUseCase(airQualityRepository) // Ajustado UseCase
	return controllers.NewGetMinutesRecordsController(getMinutesRecordService)               // Ajustado Controller
}
