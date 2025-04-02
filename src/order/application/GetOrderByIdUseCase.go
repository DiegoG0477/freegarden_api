package application

import (
	"api-order/src/order/domain/entities"
	"api-order/src/order/domain/ports"
)

type GetOrderByIdUseCase struct {
	OrderRepository ports.IOrder
}

func NewGetOrderByIdUseCase(orderRepository ports.IOrder) *GetOrderByIdUseCase {
	return &GetOrderByIdUseCase{OrderRepository: orderRepository}
}

func (o *GetOrderByIdUseCase) Run(id int64) (entities.Order, error) {
	order, err := o.OrderRepository.GetById(id)

	if err != nil {
		return entities.Order{}, err
	}

	return order, nil
}