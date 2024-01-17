package database

import (
	"fmt"
	"go-test/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// connectionString := "user=username password=password host=db sslmode=disable"
	// connectionString := "postgresql://username:password@localhost:5432/library?sslmode=disable"
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func Migrate() {
	fmt.Println("Migrating...")
	db := Init()
	db.AutoMigrate(&models.Guess{}, &models.GameType{})
}
