package ports

import "api-order/src/product/domain/entities"

type IProduct interface {
	Create(product entities.Product) (entities.Product, error)
	GetAll()([]entities.Product, error)
}