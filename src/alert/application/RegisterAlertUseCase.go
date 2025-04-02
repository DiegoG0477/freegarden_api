package application

import (
	"api-order/src/alert/domain/entities" // Adjusted import path
	"api-order/src/alert/domain/ports"    // Adjusted import path
	"errors"                              // For custom validation errors
)

type RegisterAlertUseCase struct {
	AlertRepository ports.IAlert
}

func NewRegisterAlertUseCase(alertRepo ports.IAlert) *RegisterAlertUseCase {
	return &RegisterAlertUseCase{AlertRepository: alertRepo}
}

// Run executes the logic to register a new alert
// Takes kitID, alertType, and message as input
func (uc *RegisterAlertUseCase) Run(kitID int, alertType string, message string) (entities.Alert, error) {
	// Validate alert type against known constants
	if !entities.IsValidAlertType(alertType) {
		return entities.Alert{}, errors.New("invalid alert_type provided")
	}

	alert := entities.Alert{
		KitID:     kitID,
		AlertType: alertType,
		Message:   message,
		// Timestamp will be set by the database default
	}

	createdAlert, err := uc.AlertRepository.Create(alert)
	if err != nil {
		// Handle potential errors like invalid kit_id
		return entities.Alert{}, err
	}

	return createdAlert, nil
}
