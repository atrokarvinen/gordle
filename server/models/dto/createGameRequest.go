package dto

type CreateGameRequest struct {
	Difficulty  string `json:"difficulty"`
	Language    string `json:"language"`
	MaxAttempts int    `json:"maxAttempts"`
	WordLength  int    `json:"wordLength"`
}

var DefaultCreateGameRequest CreateGameRequest = CreateGameRequest{
	Difficulty:  "all_words",
	Language:    "en",
	MaxAttempts: 6,
	WordLength:  6,
}
