package application

import (
	"api-order/src/kit/domain/entities"
	"api-order/src/kit/domain/ports"
)

type CreateKitUseCase struct {
	KitRepository ports.IKit
}

func NewCreateKitUseCase(kitRepository ports.IKit) *CreateKitUseCase {
	return &CreateKitUseCase{KitRepository: kitRepository}
}

// Run now takes userID from the controller (which gets it from JWT)
func (uc *CreateKitUseCase) Run(name, description string, userID int64) (entities.Kit, error) {
	kit := entities.Kit{
		UserID:      userID,
		Name:        name,
		Description: description,
		// CreatedAt is handled by the database DEFAULT
	}

	createdKit, err := uc.KitRepository.Create(kit)
	if err != nil {
		// Specific error handling (like unique constraints) might be needed here
		// depending on repository implementation feedback, but keep it simple for now.
		return entities.Kit{}, err
	}

	return createdKit, nil
}
