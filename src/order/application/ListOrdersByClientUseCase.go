package application

import (
	"api-order/src/order/domain/entities"
	"api-order/src/order/domain/ports"
)

type ListOrdersByClientUseCase struct {
	OrderRepository ports.IOrder
}

func NewListOrderByClientUseCase(orderRepository ports.IOrder) *ListOrdersByClientUseCase {
	return &ListOrdersByClientUseCase{OrderRepository: orderRepository}
}

func (o *ListOrdersByClientUseCase) Run(client_id int64) ([]entities.Order, error) {
	orders, err := o.OrderRepository.ListOrdersByClient(client_id)

	if err != nil {
		return []entities.Order{}, err
	}

	return orders, nil
}