package http

import (
	"api-order/src/alert/application"                     // Adjusted import path
	"api-order/src/alert/domain/ports"                    // Adjusted import path
	"api-order/src/alert/infrastructure/adapters"         // Adjusted import path
	"api-order/src/alert/infrastructure/http/controllers" // Adjusted import path
	"log"
)

// Declare repository variable specific to alerts
var (
	alertRepository ports.IAlert
)

// Initialize alert dependencies
func InitializeAlertDependencies() {
	var err error
	alertRepository, err = adapters.NewAlertRepositoryMysql()
	if err != nil {
		log.Fatalf("Error initializing alert repository: %v", err)
	}
}

// Setup function for RegisterAlertController
func SetUpRegisterAlertController() *controllers.RegisterAlertController {
	if alertRepository == nil {
		InitializeAlertDependencies()
	}
	registerAlertService := application.NewRegisterAlertUseCase(alertRepository)
	return controllers.NewRegisterAlertController(registerAlertService)
}

// Setup function for GetAlertsByKitIDController
func SetUpGetAlertsByKitIDController() *controllers.GetAlertsByKitIDController {
	if alertRepository == nil {
		InitializeAlertDependencies()
	}
	getAlertsService := application.NewGetAlertsByKitIDUseCase(alertRepository)
	return controllers.NewGetAlertsByKitIDController(getAlertsService)
}
