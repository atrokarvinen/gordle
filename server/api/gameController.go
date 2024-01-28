package api

import (
	"fmt"
	"go-test/models"
	"go-test/models/dto"
	"go-test/wordsApi"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (a Api) GetGame(c *gin.Context) {
	gameId := getIdFromParam(c)
	userId, err := getUserIdFromCookie(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}
	fmt.Printf("Getting game '%d'...\n", gameId)
	game, err := a.Game.GetGame(gameId)
	if game.UserId != userId {
		c.JSON(http.StatusForbidden, gin.H{"message": "User has no access to this game"})
		return
	}
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "game not found"})
		return
	}
	gameover := a.Game.CheckGameOver(gameId)
	dto := game
	dto.Gameover = gameover
	c.JSON(http.StatusOK, dto)
}

func (a Api) GetLatestGame(c *gin.Context) {
	fmt.Println("Getting latest game...")
	userId, err := getUserIdFromCookie(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}
	game, err := a.Game.GetLatestGame(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	gameover := a.Game.CheckGameOver(game.Id)
	dto := game
	dto.Gameover = gameover
	c.JSON(http.StatusOK, dto)
}

func (a Api) GetGames(c *gin.Context) {
	fmt.Println("Getting games...")
	games := a.Game.GetGames()
	c.JSON(http.StatusOK, games)
}

func (a Api) CreateGame(c *gin.Context) {
	fmt.Printf("Creating game...\n")
	userId, err := getUserIdFromCookie(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}
	createdGame := a.Game.CreateGame(userId)
	fmt.Println("Created game:", createdGame)
	c.JSON(http.StatusOK, createdGame)
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
	wordDetails, getWordErr := a.WordsApi.GetWord(guess.Word)
	if getWordErr != nil {
		fmt.Println("Error getting word:", getWordErr)
	}
	if getWordErr != nil && getWordErr.Error() == fmt.Sprintf("Word '%s' not found", guess.Word) {
		c.JSON(http.StatusBadRequest, gin.H{"message": getWordErr.Error()})
		return
	}
	results := a.Game.GuessWord(gameId, guess.Word)
	gameover := a.Game.CheckGameOver(gameId)
	if gameover.IsGameover {
		answerDetails := wordDetails
		if !gameover.Win || getWordErr != nil {
			answerDetails, getWordErr = a.WordsApi.GetWord(gameover.Answer)
			if getWordErr != nil {
				answerDetails = wordsApi.GetDefaultWordDetails(gameover.Answer)
			}
		}
		gameover.Answer = answerDetails.Word
		gameover.AnswerDescription = answerDetails.Results[0].Definition
	} else {
		gameover.Answer = ""
		gameover.AnswerDescription = ""
	}
	dto := dto.GuessResponse{Word: guess.Word, Results: results, Gameover: gameover}
	c.JSON(http.StatusOK, dto)
}

func getIdFromParam(c *gin.Context) int {
	gameIdStr := c.Param("id")
	gameId, _ := strconv.Atoi(gameIdStr)
	return gameId
}

func getUserIdFromCookie(c *gin.Context) (int, error) {
	userIdStr, err := c.Cookie(userIdCookie)
	if err != nil {
		return 0, err
	}
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return 0, err
	}
	return userId, nil
}
