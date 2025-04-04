package application

import (
	"api-order/src/gardendata/domain/entities" // Corrected path
	"api-order/src/gardendata/domain/ports"    // Corrected path
	"errors"
	"fmt"
)

type GetMinutesGardenDataUseCase struct {
	GardenDataRepository ports.IGardenData
}

func NewGetMinutesGardenDataUseCase(repo ports.IGardenData) *GetMinutesGardenDataUseCase {
	return &GetMinutesGardenDataUseCase{GardenDataRepository: repo}
}

// Run executes the logic to retrieve garden data records within a time window.
func (uc *GetMinutesGardenDataUseCase) Run(kitID int64, minutes int) ([]entities.GardenData, error) {
	// Basic validation
	if minutes <= 0 {
		return nil, errors.New("minutes parameter must be positive")
	}
	if kitID <= 0 {
		return nil, errors.New("kitID parameter must be positive")
	}

	records, err := uc.GardenDataRepository.GetRecordsByKitIDAndTime(kitID, minutes)
	if err != nil {
		// Log internal error details if necessary
		fmt.Printf("Error calling repository GetRecordsByKitIDAndTime for GardenData: %v\n", err)
		return nil, fmt.Errorf("failed to retrieve garden data: %w", err)
	}

	// Return empty slice if no records found (this is valid)
	if records == nil {
		return []entities.GardenData{}, nil
	}

	return records, nil
}
