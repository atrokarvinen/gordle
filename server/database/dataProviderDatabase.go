package database

import (
	"go-test/models"

	"gorm.io/gorm"
)

type DatabaseDataProvider struct {
	Db *gorm.DB
}

func (d DatabaseDataProvider) GetGame(gameId int) (models.GameType, error) {
	db := d.Db
	var game models.GameType
	result := db.First(&game, gameId)
	return game, result.Error
}

func (d DatabaseDataProvider) GetLatestGame() models.GameType {
	db := d.Db
	var game models.GameType
	db.Last(&game)
	return game
}

func (d DatabaseDataProvider) GetGames() []models.GameType {
	db := d.Db
	var games []models.GameType
	db.Find(&games)
	return games
}

func (d DatabaseDataProvider) CreateGame(game models.GameType) models.GameType {
	db := d.Db
	db.Create(&game)
	return game
}

func (d DatabaseDataProvider) GetPreviousGuesses(gameId int) []models.Guess {
	db := d.Db
	var guesses []models.Guess
	db.Where(&models.Guess{GameId: gameId}).Find(&guesses)
	return guesses
}

func (d DatabaseDataProvider) AddGuess(guess models.Guess) models.Guess {
	db := d.Db
	db.Create(&guess)
	return guess
}
