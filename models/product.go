package models

type Product struct {
	ProductId   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Category    string `json:"category"`
	Price       string `json:"price"`
	Description string `json:"description"`
}
