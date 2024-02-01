package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getIdsFromContext(c *gin.Context) (int, int, error) {
	userId, err := getUserIdFromCookie(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return 0, 0, err
	}
	gameId := getGameIdFromParam(c)
	return userId, gameId, nil
}

func getGameIdFromParam(c *gin.Context) int {
	gameIdStr := c.Param("id")
	gameId, _ := strconv.Atoi(gameIdStr)
	return gameId
}

func getUserIdFromCookie(c *gin.Context) (int, error) {
	userIdStr, err := c.Cookie(userIdCookie)
	if err != nil {
		fmt.Println("Error getting user id from cookie:", err.Error())
		return 0, err
	}
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		fmt.Println("Error converting cookie value '", userIdStr, "' to int:", err.Error())
		return 0, err
	}
	return userId, nil
}
