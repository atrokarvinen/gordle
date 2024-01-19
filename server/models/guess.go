package models

type Guess struct {
	Id     int    `json:"id"`
	GameId int    `json:"gameId"`
	Word   string `json:"word"`
}

type GuessDto struct {
	Word    string      `json:"word"`
	Letters []LetterDto `json:"letters"`
}

type LetterDto struct {
	Letter string `json:"letter"`
	State  int    `json:"state"`
}
