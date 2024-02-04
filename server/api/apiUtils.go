package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func getIdsFromContext(c *gin.Context) (int, int, error) {
	userId, err := getUserIdFromContext(c)
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

// Check cookie first, fallback to authorization header
// since some users may have 3rd party cookies disabled.
func getUserIdFromContext(c *gin.Context) (int, error) {
	userId, err := getUserIdFromCookie(c)
	if err == nil {
		return userId, nil
	}
	userId, err = getUserIdFromAuthHeader(c)
	if err == nil {
		return userId, nil
	}
	c.JSON(http.StatusForbidden, gin.H{"message": "No user id in cookie or Authorization header"})
	return 0, err
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

func getUserIdFromAuthHeader(c *gin.Context) (int, error) {
	userIdHeader := c.GetHeader("Authorization")
	if userIdHeader == "" {
		return 0, fmt.Errorf("No user id in Authorization header")
	}
	userIdStr := strings.TrimPrefix(userIdHeader, "Bearer ")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		fmt.Println("Error converting cookie value '", userIdStr, "' to int:", err.Error())
		return 0, err
	}
	return userId, nil
}
