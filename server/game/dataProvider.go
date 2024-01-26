package game

import "go-test/models"

type DataProvider interface {
	GetGame(gameId int) (models.GameType, error)
	GetLatestGame() models.GameType
	GetGames() []models.GameType
	CreateGame(game models.GameType) models.GameType
	GetPreviousGuesses(gameId int) []models.Guess
	AddGuess(guess models.Guess) models.Guess
	GetWordsApiCalls() []models.WordsApiCall
	AddWordsApiCall(wordsApiCall models.WordsApiCall) models.WordsApiCall
	DeleteWordsApiCalls(calls []models.WordsApiCall)
}
