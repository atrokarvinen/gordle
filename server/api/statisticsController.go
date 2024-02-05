package api

import "github.com/gin-gonic/gin"

func (a Api) GetStatistics(c *gin.Context) {
	userId, err := getUserIdFromContext(c)
	if err != nil {
		return
	}
	games := a.Game.GetGames(userId)
	c.JSON(200, games)
}
