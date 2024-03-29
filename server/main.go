package main

import (
	"fmt"
	"gordle/api"
	"gordle/database"
	"gordle/dictionaryClient"
	"gordle/dictionaryClient/dictionaryApi"
	"gordle/dictionaryClient/kielitoimistoApi"
	"gordle/game"
	"gordle/user"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	InitEnv()

	db := database.Init()
	database.Migrate(db)

	dataProvider := database.DatabaseDataProvider{Db: db}
	userService := user.User{Db: dataProvider}
	gameEngine := game.Game{DataProvider: dataProvider}
	clientFi := kielitoimistoApi.KielitoimistoApiClient{}
	clientEn := dictionaryApi.DictionaryApiClient{}
	clientFactory := dictionaryClient.DictionaryClientFactory{DictionaryClientEn: clientEn, DictionaryClientFi: clientFi}

	api := api.Api{DataProvider: dataProvider, Game: gameEngine, DictionaryFactory: clientFactory, User: userService}
	api.Run()
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
