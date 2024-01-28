package database

import (
	"errors"
	"fmt"
	"go-test/models/dbModels"

	"gorm.io/gorm"
)

type DatabaseDataProvider struct {
	Db *gorm.DB
}

func (d DatabaseDataProvider) GetGame(gameId int) (dbModels.Game, error) {
	var game dbModels.Game
	result := d.Db.Model(&dbModels.Game{}).Preload("Guesses").First(&game, gameId)
	fmt.Println("game guesses", game.Guesses)
	return game, result.Error
}

func (d DatabaseDataProvider) GetLatestGame(userId int) (dbModels.Game, error) {
	var game dbModels.Game
	result := d.Db.Where(&dbModels.Game{UserID: userId}).Preload("Guesses").Last(&game)
	if result.Error != nil {
		return dbModels.Game{}, errors.New("No latest game found")
	}
	return game, nil
}

func (d DatabaseDataProvider) GetGames() []dbModels.Game {
	var games []dbModels.Game
	d.Db.Preload("Guesses").Find(&games)
	return games
}

func (d DatabaseDataProvider) CreateGame(game dbModels.Game) dbModels.Game {
	d.Db.Create(&game)
	return game
}

func (d DatabaseDataProvider) GetPreviousGuesses(gameId int) []dbModels.Guess {
	var guesses []dbModels.Guess
	d.Db.Where(&dbModels.Guess{GameID: gameId}).Find(&guesses)
	return guesses
}

func (d DatabaseDataProvider) AddGuess(guess dbModels.Guess) dbModels.Guess {
	d.Db.Create(&guess)
	return guess
}

func (d DatabaseDataProvider) GetWordsApiCalls() []dbModels.WordsApiCall {
	var calls []dbModels.WordsApiCall
	d.Db.Find(&calls)
	return calls
}

func (d DatabaseDataProvider) AddWordsApiCall(wordsApiCall dbModels.WordsApiCall) dbModels.WordsApiCall {
	d.Db.Create(&wordsApiCall)
	return wordsApiCall
}

func (d DatabaseDataProvider) DeleteWordsApiCalls(calls []dbModels.WordsApiCall) {
	d.Db.Delete(&calls)
}
