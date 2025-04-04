package helpers

import (
	"api-order/src/user/application/services" // Point to user's service interface

	"golang.org/x/crypto/bcrypt"
)

type BcryptHelper struct{}

// Ensure this returns the user's IEncrypt interface type
func NewBcryptHelper() (services.IEncrypt, error) {
	return &BcryptHelper{}, nil
}

func (b *BcryptHelper) EncryptPassword(password []byte) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPwd), nil
}

func (b *BcryptHelper) ComparePassword(hashedPwd string, password []byte) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), password)
	if err != nil {
		return err // Return the specific bcrypt error
	}
	return nil
}
