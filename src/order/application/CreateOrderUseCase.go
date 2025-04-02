package application

import (
	"api-order/src/order/application/services"
	"api-order/src/order/domain/entities"
	"api-order/src/order/domain/ports"
)

type CreateOrderUseCase struct {
	OrderRepository ports.IOrder
	ProducerService services.IOrderProducer
}

func NewCreateOrderUseCase(orderRepository ports.IOrder, producerServices services.IOrderProducer) *CreateOrderUseCase {
	return &CreateOrderUseCase{OrderRepository: orderRepository, ProducerService: producerServices}
}


func (o *CreateOrderUseCase) Run(Client_id, product_id, quantity int, Total_price float64, Status string) (entities.Order, error) {
	order := entities.Order{
		Client_id: Client_id,
		Total_price: Total_price,
		Status: Status,
		Product_id: product_id,
		Quantity: quantity,
	}

	newOrder, err := o.OrderRepository.Create(order)
	o.ProducerService.PublishOrderCreated(newOrder)

	if err != nil {
		return entities.Order{}, err
	}

	return newOrder, nil
}