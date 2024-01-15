package game

import (
	"go-test/database"
	"go-test/models"
)

type DatabaseDataProvider struct {
}

func (d DatabaseDataProvider) GetGame(gameId int) (models.GameType, error) {
	db := database.Init()
	var game models.GameType
	result := db.First(&game, gameId)
	return game, result.Error
}

func (d DatabaseDataProvider) GetGames() []models.GameType {
	db := database.Init()
	var games []models.GameType
	db.Find(&games)
	return games
}

func (d DatabaseDataProvider) CreateGame(game models.GameType) models.GameType {
	db := database.Init()
	db.Create(&game)
	return game
}

func (d DatabaseDataProvider) GetPreviousGuesses(gameId int) []models.Guess {
	db := database.Init()
	var guesses []models.Guess
	db.Where(&models.Guess{GameId: gameId}).Find(&guesses)
	return guesses
}

func (d DatabaseDataProvider) AddGuess(guess models.Guess) models.Guess {
	db := database.Init()
	db.Create(&guess)
	return guess
}
