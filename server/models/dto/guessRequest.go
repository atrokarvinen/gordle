package dto

type GuessRequest struct {
	Word    string   `json:"word"`
	Letters []Letter `json:"letters"`
}

type Letter struct {
	Letter string `json:"letter"`
	State  int    `json:"state"`
}
