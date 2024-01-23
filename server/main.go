package main

import (
	"fmt"
	"go-test/api"
	"go-test/cli"
	"go-test/database"
	"go-test/game"
	"go-test/wordsApi"

	"github.com/joho/godotenv"
)

func main() {
	InitEnv()

	database.Init()
	database.Migrate()

	dataProvider := database.DatabaseDataProvider{}
	gameEngine := game.Game{DataProvider: dataProvider}
	wordsClient := wordsApi.WordsApiClient{}
	if false {
		cli := cli.Cli{DataProvider: dataProvider, Game: gameEngine}
		cli.Run()
	} else {
		api := api.Api{DataProvider: dataProvider, Game: gameEngine, WordsApi: wordsClient}
		api.Run()
	}
}

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
