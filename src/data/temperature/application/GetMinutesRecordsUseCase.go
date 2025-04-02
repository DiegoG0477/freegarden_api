package application

import (
	"api-order/src/data/temperature/domain/entities"
	"api-order/src/data/temperature/domain/ports"
	"errors" // For basic validation
)

type GetMinutesRecordsUseCase struct {
	TemperatureRepository ports.ITemperatureData
}

func NewGetMinutesRecordsUseCase(tempRepo ports.ITemperatureData) *GetMinutesRecordsUseCase {
	return &GetMinutesRecordsUseCase{TemperatureRepository: tempRepo}
}

// Run executes the logic to retrieve temperature records within a time window
func (uc *GetMinutesRecordsUseCase) Run(kitID int, minutes int) ([]entities.TemperatureData, error) {
	// Basic validation
	if minutes <= 0 {
		return nil, errors.New("minutes parameter must be positive")
	}
	if kitID <= 0 {
		return nil, errors.New("kitID parameter must be positive")
	}

	records, err := uc.TemperatureRepository.GetRecordsByKitIDAndTime(kitID, minutes)
	if err != nil {
		// Handle potential repository errors
		return nil, err
	}

	// Return empty slice if no records found, which is valid
	if records == nil {
		return []entities.TemperatureData{}, nil
	}

	return records, nil
}
