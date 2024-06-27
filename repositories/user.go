package repositories

import (
	"context"
	"encoding/json"
	"log"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) RegisterUser(ctx context.Context, params map[string]interface{}) (string, error) {
	var userID string

	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return userID, err
	}

	query := `SELECT account.register_user($1::jsonb)`

	row := r.DB.Raw(query, string(paramsJSON)).Row()
	if err := row.Scan(&userID); err != nil {
		return userID, err
	}

	return userID, nil
}

func (r *UserRepository) GetAllUsers(ctx context.Context, params map[string]interface{}) ([]map[string]interface{}, error) {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		log.Println("Error marshaling:", err)
		return nil, err
	}

	var result []byte
	query := `SELECT account.get_all_users($1::jsonb)`

	row := r.DB.Raw(query, string(paramsJSON)).Row()
	if err := row.Scan(&result); err != nil {
		log.Println("Error executing get_all_users:", err)
		return nil, err
	}

	var users []map[string]interface{}
	if err := json.Unmarshal(result, &users); err != nil {
		log.Println("Error unmarshaling:", err)
		return nil, err
	}

	return users, nil
}
