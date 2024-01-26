package database

import (
	"fmt"
	"go-test/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	connectionString := os.Getenv("DATABASE_URL")
	provider := os.Getenv("DATABASE_PROVIDER")
	var db *gorm.DB
	var err error
	if provider == "sqlite" {
		db, err = gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	} else if provider == "postgres" {
		db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	}

	if err != nil {
		panic(err)
	}

	return db
}

func Migrate(db *gorm.DB) {
	fmt.Println("Migrating...")
	db.AutoMigrate(&models.Guess{}, &models.GameType{}, &models.WordsApiCall{})
}
