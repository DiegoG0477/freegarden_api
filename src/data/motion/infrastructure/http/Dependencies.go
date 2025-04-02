package http

import (
	"api-order/src/data/motion/application"
	"api-order/src/data/motion/domain/ports"
	"api-order/src/data/motion/infrastructure/adapters"
	"api-order/src/data/motion/infrastructure/http/controllers"
	"log"
)

var (
	motionRepository ports.IMotionData
)

func InitializeMotionDependencies() {
	var err error
	if motionRepository == nil {
		motionRepository, err = adapters.NewMotionDataRepositoryMysql()
		if err != nil {
			log.Fatalf("Error initializing motion data repository: %v", err)
		}
	}
}

// Setup for RegisterRecordController (unchanged)
func SetUpRegisterRecordController() *controllers.RegisterRecordController {
	if motionRepository == nil {
		InitializeMotionDependencies()
	}
	registerRecordService := application.NewRegisterRecordUseCase(motionRepository)
	return controllers.NewRegisterRecordController(registerRecordService)
}

// Setup function for GetMinutesRecordsController <- NUEVO SETUP
func SetUpGetMinutesRecordsController() *controllers.GetMinutesRecordsController {
	if motionRepository == nil {
		InitializeMotionDependencies()
	}
	getMinutesRecordService := application.NewGetMinutesRecordsUseCase(motionRepository) // Ajustado UseCase
	return controllers.NewGetMinutesRecordsController(getMinutesRecordService)           // Ajustado Controller
}
