package models

type Game struct {
	Id          int      `json:"id"`
	MaxAttempts int      `json:"maxAttempts"`
	WordLength  int      `json:"wordLength"`
	Gameover    Gameover `json:"gameover"`
	Guesses     []Guess  `json:"guesses"`
	UserId      int      `json:"userId"`
}
