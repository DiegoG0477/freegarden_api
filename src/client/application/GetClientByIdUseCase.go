package application

import "api-order/src/client/domain/ports"

type GetClientByIdUseCase struct {
	ClientRepository ports.IClient
}

func NewGetClientByIdUseCase(clientRepository ports.IClient) *GetClientByIdUseCase {
	return &GetClientByIdUseCase{ClientRepository: clientRepository}
}

func (c *GetClientByIdUseCase) Run(id int64) (interface{}, error) {
	client, err := c.ClientRepository.GetById(id)

	if err != nil {
		return nil, err
	}

	return client, nil
}