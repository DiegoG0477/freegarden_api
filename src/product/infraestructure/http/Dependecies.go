package http

import (
	"api-order/src/product/application"
	"api-order/src/product/domain/ports"
	"api-order/src/product/infraestructure/adapters"
	"api-order/src/product/infraestructure/http/controllers"
	"log"
)

var ProductRepository ports.IProduct

func init(){
	var err error
	ProductRepository, err = adapters.NewProductRepositoryMysql()
	if err != nil {
		log.Fatalf("Error al inicializar el repositorio de productos: %s", err)
	} 
}


func SetUpCreateProduct() *controllers.CreateProductController {
	productService := application.NewCreateProductUseCase(ProductRepository)
	return controllers.NewCreateProductController(productService)
}

func SetUpGetAllProducts() *controllers.GetAllProductsController {
	productService := application.NewGetAllProductsUseCase(ProductRepository)
	return controllers.NewGetAllProductsController(productService)
}