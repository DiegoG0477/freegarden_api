package services

type IEncrypt interface {
	EncryptPassword(pwd []byte) (string, error)
	ComparePassword(hashedPwd string, plainPwd []byte) error
}
