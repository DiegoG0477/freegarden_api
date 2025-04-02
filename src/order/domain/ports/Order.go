package ports

import "api-order/src/order/domain/entities"

type IOrder interface {
	Create(order entities.Order) (entities.Order, error)
	UpdateStatus(id int64, status string) (entities.Order, error)
	GetById(id int64) (entities.Order, error)
	ListOrdersByClient(client_id int64) ([]entities.Order, error)
	Delete(id int64) (bool, error)
}