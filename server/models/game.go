package models

import "time"

type Game struct {
	Id                int       `json:"id"`
	Answer            string    `json:"-"`
	AnswerDescription []string  `json:"-"`
	AnswerExamples    []string  `json:"-"`
	Difficulty        string    `json:"difficulty"`
	MaxAttempts       int       `json:"maxAttempts"`
	WordLength        int       `json:"wordLength"`
	Language          string    `json:"language"`
	Gameover          Gameover  `json:"gameover"`
	Guesses           []Guess   `json:"guesses"`
	UserId            int       `json:"userId"`
	State             int       `json:"state"`
	CreatedAt         time.Time `json:"createdAt"`
}
