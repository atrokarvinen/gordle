package models

type Gameover struct {
	IsGameover        bool   `json:"isGameover"`
	Win               bool   `json:"win"`
	Answer            string `json:"answer"`
	AnswerDescription string `json:"answerDescription"`
}
