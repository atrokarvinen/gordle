package game

import (
	"errors"
	"fmt"
	"go-test/models"
)

type InMemoryDataProvider struct {
}

var games []models.GameType
var guesses []models.Guess

func (d InMemoryDataProvider) GetGame(gameId int) (models.GameType, error) {
	for _, game := range games {
		if game.Id == gameId {
			fmt.Println("Found game")
			return game, nil
		}
	}
	fmt.Println("Did not find game")
	return models.GameType{}, errors.New("Game not found")
}

func (d InMemoryDataProvider) GetGames() []models.GameType {
	return games
}

func (d InMemoryDataProvider) CreateGame(game models.GameType) models.GameType {
	newGame := models.GameType{Id: len(games) + 1, Name: game.Name, MaxAttempts: game.MaxAttempts, WordLength: game.WordLength}
	games = append(games, newGame)
	return newGame
}

func (d InMemoryDataProvider) GetPreviousGuesses(gameId int) []models.Guess {
	var fetchedGuesses []models.Guess
	for _, guess := range guesses {
		if guess.GameId == gameId {
			fetchedGuesses = append(fetchedGuesses, guess)
		}
	}
	return fetchedGuesses
}

func (d InMemoryDataProvider) AddGuess(guess models.Guess) models.Guess {
	newGuess := models.Guess{Id: len(guesses) + 1, GameId: guess.GameId, Word: guess.Word}
	guesses = append(guesses, newGuess)
	return newGuess
}
