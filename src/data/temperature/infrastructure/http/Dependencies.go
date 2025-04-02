package http

import (
	"api-order/src/data/temperature/application"
	"api-order/src/data/temperature/domain/ports"
	"api-order/src/data/temperature/infrastructure/adapters"
	"api-order/src/data/temperature/infrastructure/http/controllers"
	"log"
)

var (
	temperatureRepository ports.ITemperatureData
)

func InitializeTemperatureDependencies() {
	var err error
	// Prevent re-initialization if already done
	if temperatureRepository == nil {
		temperatureRepository, err = adapters.NewTemperatureDataRepositoryMysql()
		if err != nil {
			log.Fatalf("Error initializing temperature data repository: %v", err)
		}
	}
}

// Setup for RegisterRecordController (unchanged)
func SetUpRegisterRecordController() *controllers.RegisterRecordController {
	if temperatureRepository == nil {
		InitializeTemperatureDependencies()
	}
	registerRecordService := application.NewRegisterRecordUseCase(temperatureRepository)
	return controllers.NewRegisterRecordController(registerRecordService)
}

// Setup function for GetMinutesRecordsController <- NUEVO SETUP
func SetUpGetMinutesRecordsController() *controllers.GetMinutesRecordsController {
	if temperatureRepository == nil {
		InitializeTemperatureDependencies()
	}
	getMinutesRecordService := application.NewGetMinutesRecordsUseCase(temperatureRepository)
	return controllers.NewGetMinutesRecordsController(getMinutesRecordService)
}
