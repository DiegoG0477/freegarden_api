package application

import (
	kit "api-order/src/kit/domain/ports"
	"api-order/src/user/application/services"
	"api-order/src/user/domain/entities"
	"api-order/src/user/domain/ports"
	"errors"
	"fmt"
)

type RegisterUserUseCase struct {
	UserRepository ports.IUser
	KitRepository  kit.IKit
	EncryptService services.IEncrypt
}

func NewRegisterUserUseCase(userRepository ports.IUser, kitRepository kit.IKit, encryptService services.IEncrypt) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		UserRepository: userRepository,
		KitRepository:  kitRepository,
		EncryptService: encryptService,
	}
}

var ErrKitCodeExists = errors.New("kit code already exists")
var ErrUserEmailExists = errors.New("user email already exists")

func (uc *RegisterUserUseCase) Run(firstName, lastName, email, password, kitCode string) (entities.User, error) {
	// 1. Check if Kit Code already exists
	kitExists, err := uc.KitRepository.CheckKitNameExists(kitCode)
	if err != nil {
		// Log the error internally if needed
		fmt.Printf("Error checking kit name existence: %v\n", err)
		return entities.User{}, fmt.Errorf("failed to validate kit code: %w", err)
	}
	if kitExists {
		return entities.User{}, ErrKitCodeExists // Specific error for controller
	}

	// 2. Check if Email already exists (optional but good practice)
	emailExists, err := uc.UserRepository.CheckEmailExists(email)
	if err != nil {
		fmt.Printf("Error checking email existence: %v\n", err)
		return entities.User{}, fmt.Errorf("failed to validate email: %w", err)
	}
	if emailExists {
		return entities.User{}, ErrUserEmailExists
	}

	// 3. Encrypt Password
	hashPass, err := uc.EncryptService.EncryptPassword([]byte(password))
	if err != nil {
		// Log the error internally if needed
		fmt.Printf("Error encrypting password: %v\n", err)
		return entities.User{}, fmt.Errorf("failed to secure password: %w", err)
	}

	// 4. Create User Entity
	user := entities.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  hashPass, // Store the hashed password
	}

	// 5. Create User in Repository
	createdUser, err := uc.UserRepository.Create(user)
	if err != nil {
		// Log the error internally if needed
		fmt.Printf("Error creating user in repository: %v\n", err)
		// The repository might return a specific error for duplicates if CheckEmailExists wasn't used
		return entities.User{}, fmt.Errorf("failed to register user: %w", err)
	}

	// Note: Based on the prompt, we only *check* the kit code.
	// If you needed to *create* a kit entry after user creation, you would do it here:
	/*
	   kit := entities.Kit{
	       UserID: createdUser.ID,
	       Name:   kitCode,
	       // Description: "Default description or passed in",
	   }
	   _, err = uc.KitRepository.CreateKit(kit)
	   if err != nil {
	       // Handle kit creation failure - maybe rollback user creation? (complex transaction needed)
	       fmt.Printf("Warning: User created (ID: %d) but failed to create kit entry: %v\n", createdUser.ID, err)
	       // Depending on requirements, you might return an error here or just log it.
	   }
	*/

	return createdUser, nil
}
