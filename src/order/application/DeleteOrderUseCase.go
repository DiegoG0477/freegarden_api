package application

import "api-order/src/order/domain/ports"

type DeleteOrderUseCase struct {
	OrderRepository ports.IOrder
}

func NewDeleteOrderUseCase(orderRepository ports.IOrder) *DeleteOrderUseCase{
	return &DeleteOrderUseCase{OrderRepository: orderRepository}
}

func (o *DeleteOrderUseCase) Run(id int64) (bool, error) {
	deleted, err := o.OrderRepository.Delete(id)

	if err != nil {
		return false, err
	}

	return deleted, nil
}