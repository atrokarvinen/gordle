package dto

import "go-test/models"

type GuessResponse struct {
	Word     string          `json:"word"`
	Results  []string        `json:"results"`
	Gameover models.Gameover `json:"gameover"`
}
