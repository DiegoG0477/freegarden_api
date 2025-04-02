package http

import (
	"api-order/src/data/motion/application"                     // Adjusted import path
	"api-order/src/data/motion/domain/ports"                    // Adjusted import path
	"api-order/src/data/motion/infrastructure/adapters"         // Adjusted import path
	"api-order/src/data/motion/infrastructure/http/controllers" // Adjusted import path
	"log"
)

// Declare repository variable specific to motion data
var (
	motionRepository ports.IMotionData
)

// Initialize motion dependencies
func InitializeMotionDependencies() {
	var err error
	motionRepository, err = adapters.NewMotionDataRepositoryMysql()
	if err != nil {
		log.Fatalf("Error initializing motion data repository: %v", err)
	}
}

// Setup function for RegisterRecordController
func SetUpRegisterRecordController() *controllers.RegisterRecordController {
	// Ensure dependencies are initialized
	if motionRepository == nil {
		InitializeMotionDependencies()
	}
	registerRecordService := application.NewRegisterRecordUseCase(motionRepository)
	return controllers.NewRegisterRecordController(registerRecordService)
}
