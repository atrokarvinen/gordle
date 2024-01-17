package models

type GameType struct {
	Id          int
	Name        string
	Answer      string
	MaxAttempts int
	WordLength  int
}
