package helpers

import (
	"api-order/src/client/application/services"
	"golang.org/x/crypto/bcrypt"
)

type BcryptHelper struct{}

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

func (b *BcryptHelper) ComparePassword(hashedPwd string, password []byte)  error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), password)
	if err != nil {
		return  err
	}
	return  nil
}
