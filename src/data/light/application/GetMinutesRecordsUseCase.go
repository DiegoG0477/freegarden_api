package application

import (
	"api-order/src/data/light/domain/entities" // Ajustado
	"api-order/src/data/light/domain/ports"    // Ajustado
	"errors"
)

type GetMinutesRecordsUseCase struct {
	LightRepository ports.ILightData // Ajustado
}

func NewGetMinutesRecordsUseCase(lightRepo ports.ILightData) *GetMinutesRecordsUseCase { // Ajustado
	return &GetMinutesRecordsUseCase{LightRepository: lightRepo}
}

// Run executes the logic to retrieve light records within a time window
func (uc *GetMinutesRecordsUseCase) Run(kitID int, minutes int) ([]entities.LightData, error) { // Ajustado tipo de retorno
	// Basic validation
	if minutes <= 0 {
		return nil, errors.New("minutes parameter must be positive")
	}
	if kitID <= 0 {
		return nil, errors.New("kitID parameter must be positive")
	}

	records, err := uc.LightRepository.GetRecordsByKitIDAndTime(kitID, minutes)
	if err != nil {
		// Handle potential repository errors
		return nil, err
	}

	// Return empty slice if no records found, which is valid
	if records == nil {
		return []entities.LightData{}, nil // Ajustado tipo
	}

	return records, nil
}
