package models

type Gameover struct {
	IsGameover  bool     `json:"isGameover"`
	Win         bool     `json:"win"`
	Answer      string   `json:"answer"`
	Definitions []string `json:"definitions"`
	Examples    []string `json:"examples"`
}
