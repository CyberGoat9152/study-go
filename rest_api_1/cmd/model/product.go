package model

type Product struct {
	ID    int     `json:"product_id"`
	Name  string  `json:"product_name"`
	Price float64 `json:"product_price"`
}
