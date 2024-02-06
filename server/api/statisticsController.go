package api

import (
	"gordle/models/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (a Api) GetStatistics(c *gin.Context) {
	userId, err := getUserIdFromContext(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		return
	}
	games, totalCount := a.Game.GetGames(userId, page, limit)
	dto := dto.StatisticsResponse{
		TotalCount: totalCount,
		Games:      games,
	}
	c.JSON(200, dto)
}
