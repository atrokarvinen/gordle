package dto

import "gordle/models"

type StatisticsResponse struct {
	TotalCount int64         `json:"totalCount"`
	Games      []models.Game `json:"games"`
}
