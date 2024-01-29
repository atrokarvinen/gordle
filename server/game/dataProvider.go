package game

import (
	"go-test/models/dbModels"
)

type DataProvider interface {
	GetGame(gameId int) (dbModels.Game, error)
	GetLatestGame(userId int) (dbModels.Game, error)
	CreateGame(game dbModels.Game) dbModels.Game
	UpdateGame(game dbModels.Game) error
	GetPreviousGuesses(gameId int) []dbModels.Guess
	AddGuess(guess dbModels.Guess) dbModels.Guess
	GetWordsApiCalls() []dbModels.WordsApiCall
	AddWordsApiCall(wordsApiCall dbModels.WordsApiCall) dbModels.WordsApiCall
	DeleteWordsApiCalls(calls []dbModels.WordsApiCall)
}
