package application

import (
	"api-order/src/data/temperature/domain/entities"
	"api-order/src/data/temperature/domain/ports"
)

type RegisterRecordUseCase struct {
	TemperatureRepository ports.ITemperatureData
}

func NewRegisterRecordUseCase(tempRepo ports.ITemperatureData) *RegisterRecordUseCase {
	return &RegisterRecordUseCase{TemperatureRepository: tempRepo}
}

// Run executes the logic to register a new temperature/humidity record
// Takes kitID, temperature, and humidity directly as input
func (uc *RegisterRecordUseCase) Run(kitID int, temperature float64, humidity float64) (entities.TemperatureData, error) {
	data := entities.TemperatureData{
		KitID:       kitID,
		Temperature: temperature,
		Humidity:    humidity,
		// Timestamp will be set by the database default
	}

	createdRecord, err := uc.TemperatureRepository.Create(data)
	if err != nil {
		// Handle potential errors like invalid kit_id (foreign key constraint)
		return entities.TemperatureData{}, err
	}

	return createdRecord, nil
}
