package database

import (
	"errors"
	"gordle/models/dbModels"

	"gorm.io/gorm"
)

type DatabaseDataProvider struct {
	Db *gorm.DB
}

func (d DatabaseDataProvider) GetGame(gameId int) (dbModels.Game, error) {
	var game dbModels.Game
	result := d.Db.Model(&dbModels.Game{}).Preload("Guesses").First(&game, gameId)
	return game, result.Error
}

func (d DatabaseDataProvider) GetGames(userId int) []dbModels.Game {
	games := []dbModels.Game{}
	d.Db.Model(&dbModels.Game{}).Preload("Guesses").Where(dbModels.Game{UserID: userId}).Order("created_at desc").Find(&games)
	return games
}

func (d DatabaseDataProvider) GetGamesPaginated(userId int, page int, limit int) ([]dbModels.Game, int64) {
	games := []dbModels.Game{}
	offset := page * limit
	d.Db.Model(&dbModels.Game{}).Preload("Guesses").Where(dbModels.Game{UserID: userId}).Not(dbModels.Game{State: 1}).Order("created_at desc").Offset(offset).Limit(limit).Find(&games)

	var totalCount int64
	d.Db.Model(&dbModels.Game{}).Where(&dbModels.Game{UserID: userId}).Not(dbModels.Game{State: 1}).Count(&totalCount)
	return games, totalCount
}

func (d DatabaseDataProvider) GetLatestGame(userId int) (dbModels.Game, error) {
	var game dbModels.Game
	result := d.Db.Where(&dbModels.Game{UserID: userId}).Preload("Guesses").Last(&game)
	if result.Error != nil {
		return dbModels.Game{}, errors.New("No latest game found")
	}
	return game, nil
}

func (d DatabaseDataProvider) CreateGame(game dbModels.Game) dbModels.Game {
	d.Db.Create(&game)
	return game
}

func (d DatabaseDataProvider) UpdateGame(game dbModels.Game) error {
	d.Db.Model(&game).Updates(&game)
	return nil
}

func (d DatabaseDataProvider) GetPlayedAnswers(userId int) []string {
	var games []dbModels.Game
	d.Db.Where(&dbModels.Game{UserID: userId}).Find(&games)
	playedAnswers := []string{}
	for _, game := range games {
		playedAnswers = append(playedAnswers, game.Answer)
	}
	return playedAnswers
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
	d.Db.Unscoped().Delete(&calls)
}
