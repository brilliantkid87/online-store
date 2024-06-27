package repositories

import (
	"context"
	"encoding/json"
	"log"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) FindTransactions(ctx context.Context, params map[string]interface{}) ([]map[string]interface{}, error) {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		log.Println("Error marshaling params to JSON:", err)
		return nil, err
	}

	var result []byte
	query := `SELECT account.find_transactions_json($1::jsonb)`

	row := r.DB.Raw(query, string(paramsJSON)).Row()
	if err := row.Scan(&result); err != nil {
		log.Println("Error executing find_transactions_json:", err)
		return nil, err
	}

	var transactions []map[string]interface{}
	if err := json.Unmarshal(result, &transactions); err != nil {
		log.Println("Error unmarshaling result:", err)
		return nil, err
	}

	return transactions, nil
}
