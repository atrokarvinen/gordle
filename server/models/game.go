package models

type Game struct {
	Id                int      `json:"id"`
	Answer            string   `json:"-"`
	AnswerDescription string   `json:"-"`
	MaxAttempts       int      `json:"maxAttempts"`
	WordLength        int      `json:"wordLength"`
	Gameover          Gameover `json:"gameover"`
	Guesses           []Guess  `json:"guesses"`
	UserId            int      `json:"userId"`
	State             int      `json:"state"`
}
