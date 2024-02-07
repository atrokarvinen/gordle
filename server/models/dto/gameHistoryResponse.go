package dto

import "gordle/models"

type GameHistoryResponse struct {
	TotalCount int64         `json:"totalCount"`
	Games      []models.Game `json:"games"`
}
