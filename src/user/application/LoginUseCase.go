package application

import (
	"api-order/src/user/application/services"
	"api-order/src/user/domain/entities"
	"api-order/src/user/domain/ports"
	"fmt"
)

type LoginUseCase struct {
	UserRepository ports.IUser
	EncryptService services.IEncrypt
}

func NewLoginUseCase(userRepository ports.IUser, encryptService services.IEncrypt) *LoginUseCase {
	return &LoginUseCase{
		UserRepository: userRepository,
		EncryptService: encryptService,
	}
}

func (uc *LoginUseCase) Run(email string, password string) (entities.User, error) {
	user, err := uc.UserRepository.GetByEmail(email)
	if err != nil {
		// Error could be "not found" or DB error
		fmt.Printf("Error fetching user by email '%s': %v\n", email, err)
		return entities.User{}, err // Let controller interpret sql.ErrNoRows
	}

	// Compare password
	err = uc.EncryptService.ComparePassword(user.Password, []byte(password))
	if err != nil {
		// Password mismatch
		return entities.User{}, fmt.Errorf("invalid credentials") // Specific error for mismatch
	}

	// Login successful, return user data (controller will strip password)
	return user, nil
}
