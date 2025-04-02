package application

import (
	"api-order/src/order/domain/entities"
	"api-order/src/order/domain/ports"
)

type UpdateStatusOrderUseCase struct {
	OrderRepository ports.IOrder
}

func NewUpdateStatusOrderUseCase(orderRepository ports.IOrder) *UpdateStatusOrderUseCase {
	return &UpdateStatusOrderUseCase{OrderRepository: orderRepository}
}

func (o *UpdateStatusOrderUseCase) Run(id int64, status string) (entities.Order, error) {
	updated, err := o.OrderRepository.UpdateStatus(id, status)

	if err != nil {
		return entities.Order{}, nil
	}

	return updated, nil
}