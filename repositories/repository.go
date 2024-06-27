package repositories

import (

	// "log"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}
