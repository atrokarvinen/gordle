package game

import (
	"go-test/models"
	"go-test/models/dbModels"
)

func (g Game) MapDbGamesToGames(dbGames []dbModels.Game) []models.Game {
	var games []models.Game
	for _, dbGame := range dbGames {
		games = append(games, g.MapDbGameToGame(dbGame))
	}
	return games
}

func (g Game) MapDbGameToGame(dbGame dbModels.Game) models.Game {
	return models.Game{
		Id:          int(dbGame.ID),
		Guesses:     g.MapDbGuessesToGuesses(dbGame.Guesses, dbGame.Answer),
		MaxAttempts: dbGame.MaxAttempts,
		WordLength:  dbGame.WordLength,
	}
}

func (g Game) MapDbGuessesToGuesses(dbGuesses []dbModels.Guess, answer string) []models.Guess {
	var guesses []models.Guess
	for _, dbGuess := range dbGuesses {
		guesses = append(guesses, g.MapDbGuessToGuess(dbGuess, answer))
	}
	return guesses
}

func (g Game) MapDbGuessToGuess(dbGuess dbModels.Guess, answer string) models.Guess {
	results := g.CheckWord(dbGuess.Word, answer)
	letters := []models.Letter{}
	for i, result := range results {
		letters = append(letters, models.Letter{
			Letter: string(dbGuess.Word[i]),
			State:  convertLetterState(result),
		})
	}
	return models.Guess{
		Id:      int(dbGuess.ID),
		Word:    dbGuess.Word,
		GameId:  int(dbGuess.GameID),
		Letters: letters,
	}
}

func convertLetterState(result string) int {
	if result == "v" {
		return 0
	} else if result == "?" {
		return 1
	}
	return 2
}
