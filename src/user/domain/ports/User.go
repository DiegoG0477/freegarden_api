package ports

import "api-order/src/user/domain/entities"

type IUser interface {
	Create(user entities.User) (entities.User, error)
	GetById(id int64) (entities.User, error)
	GetByEmail(email string) (entities.User, error)
	Update(id int64, user entities.User) (entities.User, error)
	CheckEmailExists(email string) (bool, error) // Helper for registration check
}
