package application

import (
	"api-order/src/data/motion/domain/entities" // Adjusted import path
	"api-order/src/data/motion/domain/ports"    // Adjusted import path
)

type RegisterRecordUseCase struct {
	MotionRepository ports.IMotionData
}

func NewRegisterRecordUseCase(motionRepo ports.IMotionData) *RegisterRecordUseCase {
	return &RegisterRecordUseCase{MotionRepository: motionRepo}
}

// Run executes the logic to register a new motion detection record
// Takes kitID and motionDetected directly as input
func (uc *RegisterRecordUseCase) Run(kitID int, motionDetected bool) (entities.MotionData, error) {
	data := entities.MotionData{
		KitID:          kitID,
		MotionDetected: motionDetected,
		// Timestamp will be set by the database default
	}

	createdRecord, err := uc.MotionRepository.Create(data)
	if err != nil {
		// Handle potential errors like invalid kit_id
		return entities.MotionData{}, err
	}

	return createdRecord, nil
}
