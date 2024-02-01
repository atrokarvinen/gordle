package api

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type SetAnswerRequest struct {
	Answer string `json:"answer"`
}

func (a Api) SetAnswer(c *gin.Context) {
	err := verifyTestEnv(c)
	if err != nil {
		return
	}
	answer := SetAnswerRequest{}
	c.Bind(&answer)
	fmt.Println("Setting answer to '" + answer.Answer + "'...")
	userId, _ := getUserIdFromCookie(c)
	game, _ := a.Game.GetLatestGame(userId)
	game.Answer = answer.Answer
	a.Game.UpdateGame(game)
	c.Status(http.StatusOK)
}

func verifyTestEnv(c *gin.Context) error {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "prod" {
		c.JSON(http.StatusForbidden, gin.H{"message": "Test methods are not available in production"})
		return errors.New("Not running tests in production")
	}
	return nil
}
