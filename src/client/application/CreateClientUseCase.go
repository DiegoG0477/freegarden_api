package application

import (
	"api-order/src/client/application/services"
	"api-order/src/client/domain/entities"
	"api-order/src/client/domain/ports"
)

type CreateClientUseCase struct {
	ClientRepository ports.IClient
	EncryptService   services.IEncrypt
}

func NewCreateClientUseCase(clientRepository ports.IClient, encryptService services.IEncrypt) *CreateClientUseCase {
	return &CreateClientUseCase{ClientRepository: clientRepository, EncryptService: encryptService}
}

func (c *CreateClientUseCase) Run(Name, Email, Password, LastName string) (entities.Client, error) {
	hashPass, err := c.EncryptService.EncryptPassword([]byte(Password))

	if err != nil {
		return entities.Client{}, err
	}

	client := entities.Client{
		Name:     Name,
		LastName: LastName,
		Email:    Email,
		Password: hashPass,
	}

	created, err := c.ClientRepository.Create(client)

	if err != nil {
		return entities.Client{}, err
	}

	return created, nil
}
