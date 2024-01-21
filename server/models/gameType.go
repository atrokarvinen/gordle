package models

type GameType struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Answer      string `json:"-"`
	MaxAttempts int    `json:"maxAttempts"`
	WordLength  int    `json:"wordLength"`
}

type GameDto struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	MaxAttempts int      `json:"maxAttempts"`
	WordLength  int      `json:"wordLength"`
	Gameover    Gameover `json:"gameover"`
}
