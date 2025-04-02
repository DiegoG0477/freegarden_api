package request

type CreateClientRequest struct {
	Name  		string `json:"name" validate:"required"`
	Email     	string `json:"email" validate:"required,email"`
	Password 	string `json:"password" validate:"required,min=6"`
	Address 	string `json:"address"  validate:"required"`
}

type AuthRequest struct {
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}