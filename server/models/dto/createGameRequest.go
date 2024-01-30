package dto

type CreateGameRequest struct {
	MaxAttempts int `json:"maxAttempts"`
	WordLength  int `json:"wordLength"`
}

var DefaultCreateGameRequest CreateGameRequest = CreateGameRequest{
	MaxAttempts: 6,
	WordLength:  6,
}
