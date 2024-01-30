package dto

type CreateGameRequest struct {
	Language    string `json:"language"`
	MaxAttempts int    `json:"maxAttempts"`
	WordLength  int    `json:"wordLength"`
}

var DefaultCreateGameRequest CreateGameRequest = CreateGameRequest{
	Language:    "en",
	MaxAttempts: 6,
	WordLength:  6,
}
