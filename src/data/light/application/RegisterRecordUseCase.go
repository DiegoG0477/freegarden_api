package application

import (
	"api-order/src/data/light/domain/entities" // Adjusted import path
	"api-order/src/data/light/domain/ports"    // Adjusted import path
)

type RegisterRecordUseCase struct {
	LightRepository ports.ILightData
}

func NewRegisterRecordUseCase(lightRepo ports.ILightData) *RegisterRecordUseCase {
	return &RegisterRecordUseCase{LightRepository: lightRepo}
}

// Run executes the logic to register a new light level record
// Takes kitID and lightLevel directly as input
func (uc *RegisterRecordUseCase) Run(kitID int, lightLevel float64) (entities.LightData, error) {
	data := entities.LightData{
		KitID:      kitID,
		LightLevel: lightLevel,
		// Timestamp will be set by the database default
	}

	createdRecord, err := uc.LightRepository.Create(data)
	if err != nil {
		// Handle potential errors like invalid kit_id
		return entities.LightData{}, err
	}

	return createdRecord, nil
}
