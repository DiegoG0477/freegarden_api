package http

import (
	"api-order/src/client/application"
	"api-order/src/client/application/services"
	"api-order/src/client/domain/ports"
	"api-order/src/client/infraestructure/adapters"
	"api-order/src/client/infraestructure/http/controllers"
	"api-order/src/client/infraestructure/http/controllers/helpers"
	"log"
)

var (
	clientRepository ports.IClient
	encryptServices services.IEncrypt
)

func init(){
	var err error
	clientRepository, err = adapters.NewClientRepositoryMysql()
	if err != nil {
		log.Fatalf("Error initializing client repository: %v", err )
	}

	encryptServices, err = helpers.NewBcryptHelper()
	if err != nil {
		log.Fatalf("Error initializing encrypt service: %v", err )
	}
}

func SetUpCreateClientController() *controllers.CreateClientController {
	createService := application.NewCreateClientUseCase(clientRepository, encryptServices)
	return controllers.NewCreateClientController(createService)
}

func AuthControllers() *controllers.AuthController {
	authService := application.NewAuthUseCase(clientRepository)
	return controllers.NewAuthController(authService)
}

func GetClientByIdController() *controllers.GetClientByIdController {
	getClientService := application.NewGetClientByIdUseCase(clientRepository)
	return controllers.NewGetClientByIdController(getClientService)
}