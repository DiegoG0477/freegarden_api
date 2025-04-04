package application

import (
	"api-order/src/gardendata/domain/entities" // Corrected path
	"api-order/src/gardendata/domain/ports"    // Corrected path
	"errors"
	"fmt"
)

type RegisterGardenDataUseCase struct {
	GardenDataRepository ports.IGardenData
}

func NewRegisterGardenDataUseCase(repo ports.IGardenData) *RegisterGardenDataUseCase {
	return &RegisterGardenDataUseCase{GardenDataRepository: repo}
}

// Run executes the logic to register a new garden data record.
func (uc *RegisterGardenDataUseCase) Run(kitID int64, temperature, groundHumidity, environmentHumidity, phLevel float64, time int64) (entities.GardenData, error) {
	// Basic validation (can be expanded)
	if kitID <= 0 {
		return entities.GardenData{}, errors.New("invalid kit_id provided")
	}
	// Add other validations if needed (e.g., range checks for sensor values)

	data := entities.GardenData{
		KitID:               kitID,
		Temperature:         temperature,
		GroundHumidity:      groundHumidity,
		EnvironmentHumidity: environmentHumidity,
		PhLevel:             phLevel,
		Time:                time,
		// Timestamp will be set by the database default or repository
	}

	createdRecord, err := uc.GardenDataRepository.Create(data)
	if err != nil {
		// Log internal error details if necessary
		fmt.Printf("Error calling repository Create for GardenData: %v\n", err)
		// Return a generic error or the specific repository error if safe
		return entities.GardenData{}, fmt.Errorf("failed to register garden data: %w", err)
	}

	return createdRecord, nil
}
