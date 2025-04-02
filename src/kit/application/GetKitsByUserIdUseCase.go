package application

import (
	"api-order/src/kit/domain/entities"
	"api-order/src/kit/domain/ports"
)

type GetKitsUseCase struct {
	KitRepository ports.IKit
}

func NewGetKitsUseCase(kitRepository ports.IKit) *GetKitsUseCase {
	return &GetKitsUseCase{KitRepository: kitRepository}
}

// Run takes the userID to fetch kits for
func (uc *GetKitsUseCase) Run(userID int64) ([]entities.Kit, error) {
	kits, err := uc.KitRepository.GetByUserID(userID)
	if err != nil {
		// Handle specific errors if needed, e.g., distinguishing "not found" from other DB errors
		return nil, err
	}

	// It's okay to return an empty slice if no kits are found
	if kits == nil {
		return []entities.Kit{}, nil
	}

	return kits, nil
}
