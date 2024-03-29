package models

type Guess struct {
	Id      int      `json:"id"`
	GameId  int      `json:"gameId"`
	Word    string   `json:"word"`
	Letters []Letter `json:"letters"`
}

type Letter struct {
	Letter string `json:"letter"`
	State  int    `json:"state"`
}
