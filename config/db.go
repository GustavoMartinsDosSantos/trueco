package config

import (
	"github.com/GustavoMartinsDosSantos/trueco/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// Migrate the schema to create the tables
	db.AutoMigrate(&models.Player{}, &models.Pair{})

	DB = db
}
