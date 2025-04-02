package http

import (
	"api-order/src/data/light/application"
	"api-order/src/data/light/domain/ports"
	"api-order/src/data/light/infrastructure/adapters"
	"api-order/src/data/light/infrastructure/http/controllers"
	"log"
)

var (
	lightRepository ports.ILightData
)

func InitializeLightDependencies() {
	var err error
	if lightRepository == nil {
		lightRepository, err = adapters.NewLightDataRepositoryMysql()
		if err != nil {
			log.Fatalf("Error initializing light data repository: %v", err)
		}
	}
}

// Setup for RegisterRecordController (unchanged)
func SetUpRegisterRecordController() *controllers.RegisterRecordController {
	if lightRepository == nil {
		InitializeLightDependencies()
	}
	registerRecordService := application.NewRegisterRecordUseCase(lightRepository)
	return controllers.NewRegisterRecordController(registerRecordService)
}

// Setup function for GetMinutesRecordsController <- NUEVO SETUP
func SetUpGetMinutesRecordsController() *controllers.GetMinutesRecordsController {
	if lightRepository == nil {
		InitializeLightDependencies()
	}
	getMinutesRecordService := application.NewGetMinutesRecordsUseCase(lightRepository) // Ajustado UseCase
	return controllers.NewGetMinutesRecordsController(getMinutesRecordService)          // Ajustado Controller
}
