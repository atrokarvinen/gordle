package api

type ApiError struct {
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}
