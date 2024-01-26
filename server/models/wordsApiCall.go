package models

type WordsApiCall struct {
	Id        int    `json:"id"`
	Word      string `json:"word"`
	CreatedAt string `json:"created_at"`
}
