package game

import "go-test/models"

type InMemoryDataProvider struct {
}

var games []models.GameType
var inMemoryGuesses []models.Guess

func (d InMemoryDataProvider) CreateGame(game models.GameType) models.GameType {
	newGame := models.GameType{Id: len(games) + 1, Name: game.Name, MaxAttempts: game.MaxAttempts, WordLength: game.WordLength}
	games = append(games, newGame)
	return newGame
}

func (d InMemoryDataProvider) GetPreviousGuesses(gameId int) []models.Guess {
	var guesses []models.Guess
	for _, guess := range inMemoryGuesses {
		if guess.GameId == gameId {
			guesses = append(guesses, guess)
		}
	}
	return guesses
}

func (d InMemoryDataProvider) AddGuess(guess models.Guess) models.Guess {
	newGuess := models.Guess{Id: len(inMemoryGuesses) + 1, GameId: guess.GameId, Word: guess.Word}
	inMemoryGuesses = append(inMemoryGuesses, newGuess)
	return newGuess
}
