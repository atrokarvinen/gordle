package main

import (
	"fmt"
	"go-test/api"
	"go-test/database"
	"go-test/game"
	"go-test/user"
	"go-test/wordsApi"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	InitEnv()

	db := database.Init()
	database.Migrate(db)

	// database.Reset(db)
	// database.PrintDb(db)

	dataProvider := database.DatabaseDataProvider{Db: db}
	userService := user.User{Db: dataProvider}
	gameEngine := game.Game{DataProvider: dataProvider}
	wordsClient := wordsApi.WordsApiClient{DataProvider: dataProvider}

	if false {
		// cli := cli.Cli{DataProvider: dataProvider, Game: gameEngine}
		// cli.Run()
	} else {
		api := api.Api{DataProvider: dataProvider, Game: gameEngine, WordsApi: wordsClient, User: userService}
		api.Run()
	}

}

func InitEnv() {
	appEnv := os.Getenv("APP_ENV")
	fmt.Println("APP_ENV: " + appEnv)

	if appEnv == "prod" {
		fmt.Println("Not loading .env file in production, variables passed in elsewhere")
		return
	}

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:" + err.Error())
	}
}
