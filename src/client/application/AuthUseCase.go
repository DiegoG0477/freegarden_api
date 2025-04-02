package application

import (
	"api-order/src/client/domain/entities"
	"api-order/src/client/domain/ports"
)

type AuthUseCase struct {
	ClientRepository ports.IClient
}

func NewAuthUseCase(clientRepository ports.IClient) *AuthUseCase {
	return &AuthUseCase{ClientRepository: clientRepository}
}

func (c *AuthUseCase) Run(email string) (entities.Client, error) {
	client, err := c.ClientRepository.GetByEmail(email)

	if err != nil {
		return entities.Client{}, err
	}

	return client, nil
}