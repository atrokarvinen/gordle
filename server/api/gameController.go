package api

import (
	"fmt"
	"go-test/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (a Api) GetGame(c *gin.Context) {
	gameId := getIdFromParam(c)
	fmt.Printf("Getting game '%d'...\n", gameId)
	game, err := a.Game.LoadGame(gameId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "game not found"})
		return
	}
	c.JSON(http.StatusOK, game)
}

func (a Api) GetLatestGame(c *gin.Context) {
	fmt.Println("Getting latest game...")
	game := a.Game.GetLatestGame()
	c.JSON(http.StatusOK, game)
}

func (a Api) GetGames(c *gin.Context) {
	fmt.Println("Getting games...")
	games := a.Game.GetGames()
	c.JSON(http.StatusOK, games)
}

func (a Api) GetGuesses(c *gin.Context) {
	gameId := getIdFromParam(c)
	fmt.Printf("Getting guesses for game '%d'...\n", gameId)
	game, _ := a.DataProvider.GetGame(gameId)
	guesses := a.Game.GetGuesses(gameId)
	guessesDto := make([]models.GuessDto, len(guesses))
	for i, guess := range guesses {
		results := a.Game.CheckWord(guess.Word, game.Answer, len(guess.Word))
		letters := make([]models.LetterDto, len(guess.Word))
		for j, letter := range guess.Word {
			state := convertLetterState(results[j])
			letters[j] = models.LetterDto{Letter: string(letter), State: state}
		}
		guessesDto[i] = models.GuessDto{Word: guess.Word, Letters: letters}
	}
	c.JSON(http.StatusOK, guessesDto)
}

func convertLetterState(result string) int {
	if result == "v" {
		return 0
	} else if result == "?" {
		return 1
	}
	return 2
}

func (a Api) CreateGame(c *gin.Context) {
	var game models.GameType = models.GameType{Name: "New Game"}
	fmt.Printf("Creating game '%s'...\n", game.Name)
	game = a.Game.CreateGame(game.Name)
	c.JSON(http.StatusOK, game)
}

func (a Api) UpdateGame(c *gin.Context) {
	gameId := getIdFromParam(c)
	fmt.Printf("Updating game '%d'...\n", gameId)
	c.Status(http.StatusOK)
}

func (a Api) DeleteGame(c *gin.Context) {
	gameId := getIdFromParam(c)
	fmt.Printf("Deleting game '%d'...\n", gameId)
	c.Status(http.StatusOK)
}

func (a Api) GuessWord(c *gin.Context) {
	fmt.Println("Guessing word...")
	gameId := getIdFromParam(c)
	var guess models.Guess
	c.BindJSON(&guess)
	fmt.Printf("Guessing word %q for game '%d'\n", guess.Word, gameId)
	results := a.Game.GuessWord(gameId, guess.Word)
	gameover := a.Game.CheckGameOver(gameId)
	var dto models.GuessResultDto = models.GuessResultDto{Word: guess.Word, Results: results, Gameover: gameover}
	c.JSON(http.StatusOK, dto)
}

func getIdFromParam(c *gin.Context) int {
	gameIdStr := c.Param("id")
	gameId, _ := strconv.Atoi(gameIdStr)
	return gameId
}
