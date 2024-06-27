package repositories

import (
	"context"
	"encoding/json"
	"log"
	"synapsis/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) RegisterProduct(ctx context.Context, params map[string]interface{}) (string, error) {
	var productId string

	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return productId, err
	}

	query := `SELECT account.create_product($1::jsonb)`

	row := r.DB.Raw(query, string(paramsJSON)).Row()
	if err := row.Scan(&productId); err != nil {
		return productId, err
	}

	log.Println("product id: ", productId)
	return productId, nil
}

// func (r *ShoppingCartRepository) CheckoutAndPay(ctx context.Context, params map[string]interface{}) (string, error) {
// 	paramsJSON, err := json.Marshal(params)
// 	if err != nil {
// 		log.Println("Error marshaling:", err)
// 		return "", err
// 	}

// 	var transactionID string
// 	query := `SELECT account.checkout_and_pay($1::jsonb)`

// 	row := r.DB.Raw(query, string(paramsJSON)).Row()
// 	if err := row.Scan(&transactionID); err != nil {
// 		log.Println("Error executing checkout_and_pay:", err)
// 		return "", err
// 	}

// 	return transactionID, nil
// }

func (r *ProductRepository) GetProductsByCategory(category string) ([]models.Product, error) {
	var products []models.Product

	// Adjust the query to call the UDF correctly
	result := r.DB.Raw("SELECT * FROM account.get_products_by_category_v2(?)", category).Scan(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
