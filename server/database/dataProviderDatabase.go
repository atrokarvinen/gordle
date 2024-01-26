package database

import (
	"go-test/models"

	"gorm.io/gorm"
)

type DatabaseDataProvider struct {
	Db *gorm.DB
}

func (d DatabaseDataProvider) GetGame(gameId int) (models.GameType, error) {
	var game models.GameType
	result := d.Db.First(&game, gameId)
	return game, result.Error
}

func (d DatabaseDataProvider) GetLatestGame() models.GameType {
	var game models.GameType
	d.Db.Last(&game)
	return game
}

func (d DatabaseDataProvider) GetGames() []models.GameType {
	var games []models.GameType
	d.Db.Find(&games)
	return games
}

func (d DatabaseDataProvider) CreateGame(game models.GameType) models.GameType {
	d.Db.Create(&game)
	return game
}

func (d DatabaseDataProvider) GetPreviousGuesses(gameId int) []models.Guess {
	var guesses []models.Guess
	d.Db.Where(&models.Guess{GameId: gameId}).Find(&guesses)
	return guesses
}

func (d DatabaseDataProvider) AddGuess(guess models.Guess) models.Guess {
	d.Db.Create(&guess)
	return guess
}

func (d DatabaseDataProvider) GetWordsApiCalls() []models.WordsApiCall {
	var calls []models.WordsApiCall
	d.Db.Find(&calls)
	return calls
}

func (d DatabaseDataProvider) AddWordsApiCall(wordsApiCall models.WordsApiCall) models.WordsApiCall {
	d.Db.Create(&wordsApiCall)
	return wordsApiCall
}

func (d DatabaseDataProvider) DeleteWordsApiCalls(calls []models.WordsApiCall) {
	d.Db.Delete(&calls)
}
