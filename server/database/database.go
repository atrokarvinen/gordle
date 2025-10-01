package database

import (
	"fmt"
	"os"

	"gordle/models/dbModels"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Init() *gorm.DB {
	connectionString := os.Getenv("DATABASE_URL")
	provider := os.Getenv("DATABASE_PROVIDER")
	var db *gorm.DB
	var err error
	fmt.Println("Using database provider: " + provider)

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

	db.AutoMigrate(&dbModels.User{})
	db.AutoMigrate(&dbModels.Game{})
	db.AutoMigrate(&dbModels.Guess{})
	db.AutoMigrate(&dbModels.WordsApiCall{})
}

func Seed(db *gorm.DB) {
	user := dbModels.User{}
	db.Create(&user)
	game := dbModels.Game{
		Answer: "answer",
		UserID: int(user.ID),
	}
	db.Create(&game)
	guess := dbModels.Guess{
		Word:   "guess",
		GameID: int(game.ID),
	}
	db.Create(&guess)
}

func Drop(db *gorm.DB) {
	db.Migrator().DropTable(
		&dbModels.Guess{},
		&dbModels.Game{},
		&dbModels.User{},
		&dbModels.WordsApiCall{})
}

func PrintDb(db *gorm.DB) {
	var games []dbModels.Game
	db.Preload("Guesses").Find(&games)
	for _, game := range games {
		fmt.Println("game id:", game.ID)
		fmt.Println("game guesses count:", len(game.Guesses))
	}
	var users []dbModels.User
	db.Preload("Games.Guesses").Preload(clause.Associations).Find(&users)
	for _, user := range users {
		fmt.Println("user id:", user.ID)
		fmt.Println("user games count:", len(user.Games))
		if len(user.Games) > 0 {
			fmt.Println("user games guesses count:", len(user.Games[0].Guesses))
		}
	}
}
