package game

import "go-test/models"

type DataProvider interface {
	CreateGame(game models.GameType) models.GameType
	GetPreviousGuesses(gameId int) []models.Guess
	AddGuess(guess models.Guess) models.Guess
}
