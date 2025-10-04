package main

import (
	"fmt"
	"gordle/api"
	"gordle/dictionaryClient/freeDictionaryApi"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	clientGeneral := freeDictionaryApi.FreeDictionaryApiClient{}
	api.VerifyWords(clientGeneral, "de", 8)

	api.VerifyWords(clientGeneral, "sv", 5)
	api.VerifyWords(clientGeneral, "sv", 6)
	api.VerifyWords(clientGeneral, "sv", 7)
	api.VerifyWords(clientGeneral, "sv", 8)

	api.VerifyWords(clientGeneral, "pl", 5)
	api.VerifyWords(clientGeneral, "pl", 6)
	api.VerifyWords(clientGeneral, "pl", 7)
	api.VerifyWords(clientGeneral, "pl", 8)

	// tools.ParseTxtAnswers()
	// InitEnv()

	// db := database.Init()
	// database.Migrate(db)

	// dataProvider := database.DatabaseDataProvider{Db: db}
	// userService := user.User{Db: dataProvider}
	// gameEngine := game.Game{DataProvider: dataProvider}
	// clientFi := kielitoimistoApi.KielitoimistoApiClient{}
	// clientEn := dictionaryApi.DictionaryApiClient{}
	// clientGeneral := freeDictionaryApi.FreeDictionaryApiClient{}
	// clientFactory := dictionaryClient.DictionaryClientFactory{
	// 	DictionaryClientEn:      clientEn,
	// 	DictionaryClientFi:      clientFi,
	// 	DictionaryClientGeneral: clientGeneral,
	// }

	// api := api.Api{DataProvider: dataProvider, Game: gameEngine, DictionaryFactory: clientFactory, User: userService}
	// api.Run()
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
