package request


type CreateOrderRequest struct {
	Client_id int `json:"client_id" validate:"required"`
	Product_id int `json:"product_id" validate:"required"`
	Quantity int `json:"quantity" validate:"required"`
	Status string `json:"status" validate:"required,oneof=Pending processing completedd cancelled"`
	Total_price float64 `json:"total_price" validate:"required"`
}

type UpdateOrderRequest struct {
	ID int `json:"id"`
	Status string `json:"status" validate:"required,oneof=Pending processing completedd cancelled"`
}