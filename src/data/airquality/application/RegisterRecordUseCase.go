package application

import (
	"api-order/src/data/airquality/domain/entities" // Adjusted import path
	"api-order/src/data/airquality/domain/ports"    // Adjusted import path
)

type RegisterRecordUseCase struct {
	AirQualityRepository ports.IAirQualityData
}

func NewRegisterRecordUseCase(aqRepo ports.IAirQualityData) *RegisterRecordUseCase {
	return &RegisterRecordUseCase{AirQualityRepository: aqRepo}
}

// Run executes the logic to register a new air quality record
// Takes kitID and airQualityIndex directly as input
func (uc *RegisterRecordUseCase) Run(kitID int, airQualityIndex int) (entities.AirQualityData, error) {
	data := entities.AirQualityData{
		KitID:           kitID,
		AirQualityIndex: airQualityIndex,
		// Timestamp will be set by the database default
	}

	createdRecord, err := uc.AirQualityRepository.Create(data)
	if err != nil {
		// Handle potential errors like invalid kit_id (foreign key constraint)
		return entities.AirQualityData{}, err
	}

	return createdRecord, nil
}
