package dbModels

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	Answer            string
	AnswerDescription string
	MaxAttempts       int
	WordLength        int
	Language          string
	UserID            int
	Guesses           []Guess
	State             int
}

type GameState int

const (
	Unknown GameState = iota
	Active
	Win
	Lose
)
