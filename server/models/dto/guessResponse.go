package dto

import "gordle/models"

type GuessResponse struct {
	Word     string          `json:"word"`
	Results  []string        `json:"results"`
	Gameover models.Gameover `json:"gameover"`
}
