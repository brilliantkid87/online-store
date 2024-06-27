package repositories

import (
	"context"
	"encoding/json"
	"log"
	"synapsis/models"

	"gorm.io/gorm"
)

type ShoppingCartRepository struct {
	DB *gorm.DB
}

func NewShoppingCartRepository(db *gorm.DB) *ShoppingCartRepository {
	return &ShoppingCartRepository{DB: db}
}

func (r *ShoppingCartRepository) AddToCart(ctx context.Context, cart *models.ShoppingCart) error {
	paramsJSON, err := json.Marshal(cart)
	if err != nil {
		log.Println("Error marshaling params:", err)
		return err
	}

	query := `SELECT account.add_to_cart($1::jsonb)`

	result := r.DB.Exec(query, string(paramsJSON))
	if result.Error != nil {
		log.Println("Error executing add_to_cart:", result.Error)
		return result.Error
	}

	return nil
}

func (r *ShoppingCartRepository) GetCartItems(ctx context.Context, userID string) (map[string]interface{}, error) {
	var result []byte

	query := `SELECT account.get_cart_items($1)`

	row := r.DB.Raw(query, userID).Row()
	if err := row.Scan(&result); err != nil {
		log.Println("Error executing get_cart_items:", err)
		return nil, err
	}

	var items map[string]interface{}
	if err := json.Unmarshal(result, &items); err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return nil, err
	}

	return items, nil
}

func (r *ShoppingCartRepository) CheckoutAndPay(ctx context.Context, params map[string]interface{}) (string, error) {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		log.Println("Error marshaling:", err)
		return "", err
	}

	var transactionID string
	query := `SELECT account.checkout_and_pay($1::jsonb)`

	row := r.DB.Raw(query, string(paramsJSON)).Row()
	if err := row.Scan(&transactionID); err != nil {
		log.Println("Error executing checkout_and_pay:", err)
		return "", err
	}

	return transactionID, nil
}

func (r *ShoppingCartRepository) DeleteProductInCart(ctx context.Context, cart *models.ShoppingCart) error {
	paramsJSON, err := json.Marshal(cart)
	if err != nil {
		log.Println("Error marshaling params:", err)
		return err
	}

	query := `SELECT account.remove_from_cart($1::jsonb)`

	result := r.DB.Exec(query, string(paramsJSON))
	if result.Error != nil {
		log.Println("Error executing remove product in cart:", result.Error)
		return result.Error
	}

	return nil
}
