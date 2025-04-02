package application

import (
	"api-order/src/data/motion/domain/entities" // Ajustado
	"api-order/src/data/motion/domain/ports"    // Ajustado
	"errors"
)

type GetMinutesRecordsUseCase struct {
	MotionRepository ports.IMotionData // Ajustado
}

func NewGetMinutesRecordsUseCase(motionRepo ports.IMotionData) *GetMinutesRecordsUseCase { // Ajustado
	return &GetMinutesRecordsUseCase{MotionRepository: motionRepo}
}

// Run executes the logic to retrieve motion records within a time window
func (uc *GetMinutesRecordsUseCase) Run(kitID int, minutes int) ([]entities.MotionData, error) { // Ajustado tipo de retorno
	// Basic validation
	if minutes <= 0 {
		return nil, errors.New("minutes parameter must be positive")
	}
	if kitID <= 0 {
		return nil, errors.New("kitID parameter must be positive")
	}

	records, err := uc.MotionRepository.GetRecordsByKitIDAndTime(kitID, minutes)
	if err != nil {
		// Handle potential repository errors
		return nil, err
	}

	// Return empty slice if no records found, which is valid
	if records == nil {
		return []entities.MotionData{}, nil // Ajustado tipo
	}

	return records, nil
}
