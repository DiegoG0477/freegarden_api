package application

import (
	"api-order/src/alert/domain/entities" // Adjusted import path
	"api-order/src/alert/domain/ports"    // Adjusted import path
)

type GetAlertsByKitIDUseCase struct {
	AlertRepository ports.IAlert
}

func NewGetAlertsByKitIDUseCase(alertRepo ports.IAlert) *GetAlertsByKitIDUseCase {
	return &GetAlertsByKitIDUseCase{AlertRepository: alertRepo}
}

// Run executes the logic to retrieve alerts for a specific kit ID
func (uc *GetAlertsByKitIDUseCase) Run(kitID int) ([]entities.Alert, error) {
	alerts, err := uc.AlertRepository.GetByKitID(kitID)
	if err != nil {
		// Handle potential errors (e.g., DB connection issues)
		return nil, err
	}

	// Return an empty slice if no alerts found, which is valid
	if alerts == nil {
		return []entities.Alert{}, nil
	}

	return alerts, nil
}
