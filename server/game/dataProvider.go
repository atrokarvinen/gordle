package game

import (
	"gordle/models/dbModels"
)

type DataProvider interface {
	GetGame(gameId int) (dbModels.Game, error)
	GetGames(userId int) []dbModels.Game
	GetLatestGame(userId int) (dbModels.Game, error)
	CreateGame(game dbModels.Game) dbModels.Game
	UpdateGame(game dbModels.Game) error
	GetPlayedAnswers(userId int) []string
	AddGuess(guess dbModels.Guess) dbModels.Guess
	GetWordsApiCalls() []dbModels.WordsApiCall
	AddWordsApiCall(wordsApiCall dbModels.WordsApiCall) dbModels.WordsApiCall
	DeleteWordsApiCalls(calls []dbModels.WordsApiCall)
}
