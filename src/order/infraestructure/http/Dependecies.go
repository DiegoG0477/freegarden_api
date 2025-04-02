package http

import (
	"api-order/src/order/application"
	"api-order/src/order/application/services"
	"api-order/src/order/domain/ports"
	"api-order/src/order/infraestructure/adapters"
	"api-order/src/order/infraestructure/http/controllers"
	"api-order/src/order/infraestructure/http/controllers/helpers"
	"log"
)

var (
	OrderRepository ports.IOrder
	ProductorService services.IOrderProducer
)

func init(){
	var err error

	OrderRepository, err = adapters.NewOrderRepositoryMysql()

	if err != nil{
		log.Fatalf("Error al inicializar el repositorio de ordenes: %s", err)
	}

	ProductorService, err = helpers.NewRabbitMQProducer("order_topic")
	if err != nil{
		log.Fatalf("Error al inicializar el productor de ordenes: %s", err)
	}
}


func SetUpCreate() *controllers.CreateOrderController {
	orderService := application.NewCreateOrderUseCase(OrderRepository, ProductorService)
	return controllers.NewCreateOrderController(orderService)
}

func ListOrdersByClientController() *controllers.ListOrdersByClientController{
	orderService := application.NewListOrderByClientUseCase(OrderRepository)
	return controllers.NewListOrderByClientController(orderService)
}



func GetByIdController() *controllers.GetOrderByIdController{
	orderService := application.NewGetOrderByIdUseCase(OrderRepository)
	return controllers.NewGetOrderByIdController(orderService)
}

func Delete() *controllers.DeleteOrderController{
	orderService := application.NewDeleteOrderUseCase(OrderRepository)
	return controllers.NewDeleteOrderController(orderService)
}

func UpdateStatusController() *controllers.UpdateStatusController{
	orderService := application.NewUpdateStatusOrderUseCase(OrderRepository)
	return controllers.NewUpdateStatusController(orderService)
}