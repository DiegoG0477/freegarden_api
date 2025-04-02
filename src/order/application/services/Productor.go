package services

import "api-order/src/order/domain/entities"

type IOrderProducer interface {
	PublishOrderCreated(order entities.Order) error
}
