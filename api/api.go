package api

import (
	"fmt"
	"go-test/game"
	"go-test/models"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Api struct {
	DataProvider game.DataProvider
	Game         game.Game
}

func (a Api) GetGame(c *gin.Context) {
	gameId := GetIdFromParam(c)
	fmt.Printf("Getting game '%d'...\n", gameId)
	game, err := a.Game.LoadGame(gameId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "game not found"})
		return
	}
	c.JSON(http.StatusOK, game)
}

func (a Api) GetGames(c *gin.Context) {
	fmt.Println("Getting games...")
	games := a.Game.GetGames()
	c.JSON(http.StatusOK, games)
}

func (a Api) CreateGame(c *gin.Context) {
	var game models.GameType
	c.BindJSON(&game)
	fmt.Printf("Creating game '%s'...\n", game.Name)
	game = a.Game.CreateGame(game.Name)
	c.JSON(http.StatusOK, game)
}

func (a Api) UpdateGame(c *gin.Context) {
	gameId := GetIdFromParam(c)
	fmt.Printf("Updating game '%d'...\n", gameId)
	c.Status(http.StatusOK)
}

func (a Api) DeleteGame(c *gin.Context) {
	gameId := GetIdFromParam(c)
	fmt.Printf("Deleting game '%d'...\n", gameId)
	c.Status(http.StatusOK)
}

func (a Api) GuessWord(c *gin.Context) {
	gameIdStr := c.Param("id")
	gameId, _ := strconv.Atoi(gameIdStr)
	var guess models.Guess
	c.BindJSON(&guess)
	results := a.Game.GuessWord(gameId, guess.Word)
	c.JSON(http.StatusOK, results)
}

func (a Api) Run() {
	mode := os.Getenv("GIN_MODE")
	gin.SetMode(mode)
	r := gin.Default()

	r.GET("/games", a.GetGames)
	r.GET("/games/:id", a.GetGame)
	r.POST("/games", a.CreateGame)
	r.PUT("/games/:id", a.UpdateGame)
	r.DELETE("/games/:id", a.DeleteGame)

	r.POST("/games/:id/guesses", a.GuessWord)

	r.Run("localhost:9000")
}

func GetIdFromParam(c *gin.Context) int {
	gameIdStr := c.Param("id")
	gameId, _ := strconv.Atoi(gameIdStr)
	return gameId
}
