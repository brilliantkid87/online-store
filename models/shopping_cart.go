package models

type ShoppingCart struct {
	CartId    string `json:"cart_id"`
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
