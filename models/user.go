package models

type User struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	UserName string `json:"user_name" validate:"required"`
}
