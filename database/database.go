package database

import (
	"fmt"
	"go-test/models"
	"os"

	"github.com/joho/godotenv"
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

	fmt.Println("Successfully connected to db!")

	return db
}

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	fmt.Println("CGO_ENABLED:", os.Getenv("CGO_ENABLED"))
}

func Migrate() {
	fmt.Println("Migrating...")
	db := Init()
	db.AutoMigrate(&models.Guess{}, &models.GameType{})
}
