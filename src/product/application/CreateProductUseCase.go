package application

import (
	"api-order/src/product/domain/entities"
	"api-order/src/product/domain/ports"
)

type CreateProductUseCase struct {
	ProductRepository ports.IProduct
}

func NewCreateProductUseCase(productRepository ports.IProduct) *CreateProductUseCase {
	return &CreateProductUseCase{ProductRepository: productRepository}
}

func (p *CreateProductUseCase) Run(Name, Description string, Price float64, Stock int) (entities.Product, error) {
	product := entities.Product{
		Name:        Name,
		Description: Description,
		Price:       Price,
		Stock: Stock,
	}

	newProduct, err := p.ProductRepository.Create(product)

	if err != nil {
		return entities.Product{}, err
	}

	return newProduct, nil
}