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

	// clientFi2 := wordsApi.DictClientFi{}
	// details, _ := clientFi2.GetWord("toiste")
	// fmt.Println("definitions:", details.Definitions)

	// database.Reset(db)
	// database.PrintDb(db)

	dataProvider := database.DatabaseDataProvider{Db: db}
	userService := user.User{Db: dataProvider}
	gameEngine := game.Game{DataProvider: dataProvider}
	clientFi := wordsApi.DictClientFi{}
	clientEn := wordsApi.WordsApiClient{DataProvider: dataProvider}
	clientFactory := wordsApi.DictionaryClientFactory{DictionaryClientEn: clientEn, DictionaryClientFi: clientFi}

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
