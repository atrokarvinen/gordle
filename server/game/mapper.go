package game

import (
	"go-test/models"
	"go-test/models/dbModels"

	"gorm.io/gorm"
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
		Id:                int(dbGame.ID),
		Guesses:           g.MapDbGuessesToGuesses(dbGame.Guesses, dbGame.Answer),
		MaxAttempts:       dbGame.MaxAttempts,
		WordLength:        dbGame.WordLength,
		Language:          dbGame.Language,
		UserId:            dbGame.UserID,
		Answer:            dbGame.Answer,
		AnswerDescription: dbGame.AnswerDescription,
		State:             dbGame.State,
		Gameover:          MapDbGameToGameover(dbGame),
	}
}

func MapDbGameToGameover(game dbModels.Game) models.Gameover {

	return models.Gameover{
		IsGameover:        game.State != int(dbModels.Active),
		Win:               game.State == int(dbModels.Win),
		Answer:            game.Answer,
		AnswerDescription: game.AnswerDescription,
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

func (g Game) MapGameToDbGame(game models.Game) dbModels.Game {
	return dbModels.Game{
		Model:             gorm.Model{ID: uint(game.Id)},
		Answer:            game.Answer,
		MaxAttempts:       game.MaxAttempts,
		WordLength:        game.WordLength,
		Language:          game.Language,
		UserID:            game.UserId,
		State:             game.State,
		AnswerDescription: game.AnswerDescription,
	}
}
