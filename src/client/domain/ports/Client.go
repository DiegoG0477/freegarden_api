package ports

import "api-order/src/client/domain/entities"

type IClient interface {
	Create(client entities.Client) (entities.Client, error)
	GetById(id int64) (entities.Client, error)
	GetByEmail(email string) (entities.Client, error)
}