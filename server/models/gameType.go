package models

type GameType struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Answer      string `json:"-"`
	MaxAttempts int    `json:"maxAttempts"`
	WordLength  int    `json:"wordLength"`
}
