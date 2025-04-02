package entities

type Order struct {
	ID int `json:"id"`
	Client_id int `json:"client_id"`
	Total_price float64 `json:"total_price"`
	Status string `json:"status"`
	Product_id int `json:"product_id"`
	Quantity int `json:"quantity"`
}