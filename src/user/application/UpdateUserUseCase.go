package application

import (
	"api-order/src/user/domain/entities"
	"api-order/src/user/domain/ports"
	"fmt"
)

type UpdateUserUseCase struct {
	UserRepository ports.IUser
}

func NewUpdateUserUseCase(userRepository ports.IUser) *UpdateUserUseCase {
	return &UpdateUserUseCase{UserRepository: userRepository}
}

// Consider adding password update logic separately if needed
func (uc *UpdateUserUseCase) Run(id int64, firstName, lastName string) (entities.User, error) {
	// You might want to fetch the user first to ensure they exist,
	// though the Update operation in the repo might handle "not found".
	// _, err := uc.UserRepository.GetById(id)
	// if err != nil {
	//     return entities.User{}, err // Propagate "not found" or other errors
	// }

	userToUpdate := entities.User{
		FirstName: firstName,
		LastName:  lastName,
		// Email is usually not updated or handled specially due to uniqueness
		// Password update would require current password verification and hashing
	}

	updatedUser, err := uc.UserRepository.Update(id, userToUpdate)
	if err != nil {
		fmt.Printf("Error updating user ID %d: %v\n", id, err)
		return entities.User{}, fmt.Errorf("failed to update user: %w", err)
	}

	return updatedUser, nil
}
