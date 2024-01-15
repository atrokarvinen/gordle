package main

import (
	"go-test/cli"
	"go-test/database"
	"go-test/game"
)

func main() {
	if true {
		database.Init()
		database.Migrate()
	}

	dataProvider := database.DatabaseDataProvider{}
	// dataProvider := game.InMemoryDataProvider{}
	gameEngine := game.Game{DataProvider: dataProvider}
	cli := cli.Cli{DataProvider: dataProvider, Game: gameEngine}

	cli.Run()
}
