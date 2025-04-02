package ports

import "api-order/src/kit/domain/entities"

type IKit interface {
	Create(kit entities.Kit) (entities.Kit, error)
	GetByUserID(userID int64) ([]entities.Kit, error)
}
