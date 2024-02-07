package api

import (
	"gordle/models/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (a Api) GetGameHistory(c *gin.Context) {
	userId, err := getUserIdFromContext(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		return
	}
	games, totalCount := a.Game.GetGames(userId, page, limit)
	dto := dto.GameHistoryResponse{
		TotalCount: totalCount,
		Games:      games,
	}
	c.JSON(http.StatusOK, dto)
}

func (a Api) GetStatistics(c *gin.Context) {
	userId, err := getUserIdFromContext(c)
	if err != nil {
		return
	}
	games := a.Game.GetStatistics(userId)
	dto := dto.StatisticsResponse{
		AllGames: games,
	}
	c.JSON(http.StatusOK, dto)
}
