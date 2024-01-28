package dbModels

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	Answer      string
	MaxAttempts int
	WordLength  int
	UserID      int
	Guesses     []Guess
}
