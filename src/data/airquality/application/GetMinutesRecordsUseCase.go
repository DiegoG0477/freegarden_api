package application

import (
	"api-order/src/data/airquality/domain/entities" // Ajustado
	"api-order/src/data/airquality/domain/ports"    // Ajustado
	"errors"
)

type GetMinutesRecordsUseCase struct {
	AirQualityRepository ports.IAirQualityData // Ajustado
}

func NewGetMinutesRecordsUseCase(aqRepo ports.IAirQualityData) *GetMinutesRecordsUseCase { // Ajustado
	return &GetMinutesRecordsUseCase{AirQualityRepository: aqRepo}
}

// Run executes the logic to retrieve air quality records within a time window
func (uc *GetMinutesRecordsUseCase) Run(kitID int, minutes int) ([]entities.AirQualityData, error) { // Ajustado tipo de retorno
	// Basic validation
	if minutes <= 0 {
		return nil, errors.New("minutes parameter must be positive")
	}
	if kitID <= 0 {
		return nil, errors.New("kitID parameter must be positive")
	}

	records, err := uc.AirQualityRepository.GetRecordsByKitIDAndTime(kitID, minutes)
	if err != nil {
		// Handle potential repository errors
		return nil, err
	}

	// Return empty slice if no records found, which is valid
	if records == nil {
		return []entities.AirQualityData{}, nil // Ajustado tipo
	}

	return records, nil
}
