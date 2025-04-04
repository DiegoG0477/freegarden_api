package application

import (
	"api-order/src/user/domain/entities"
	"api-order/src/user/domain/ports"
	"fmt"
)

type GetUserByIdUseCase struct {
	UserRepository ports.IUser
}

func NewGetUserByIdUseCase(userRepository ports.IUser) *GetUserByIdUseCase {
	return &GetUserByIdUseCase{UserRepository: userRepository}
}

func (uc *GetUserByIdUseCase) Run(id int64) (entities.User, error) {
	user, err := uc.UserRepository.GetById(id)
	if err != nil {
		fmt.Printf("Error fetching user by ID %d: %v\n", id, err)
		return entities.User{}, err // Let controller interpret sql.ErrNoRows
	}
	return user, nil
}
