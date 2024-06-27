package database

import (
	"log"
	"synapsis/models"
	"synapsis/pkg/psql"
)

func RunMigration() {
	err := psql.DB.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		panic("Migration Failed")
	}

	log.Println("Migration Success")
}
