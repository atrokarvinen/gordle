package main

import (
	"go-test/cli"
	"go-test/database"
)

func main() {
	if true {
		database.Init()
		database.Migrate()
	}
	var cli = cli.Cli{}
	cli.Run()
}
