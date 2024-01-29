package api

import (
	"fmt"
	"go-test/models"
	"go-test/models/dto"
	"go-test/wordsApi"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (a Api) GetGame(c *gin.Context) {
	gameId := getGameIdFromParam(c)
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
	userId, err := getUserIdFromCookie(c)
	fmt.Println("Getting latest game for user '", userId, "'...")
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
	userId, err := getUserIdFromCookie(c)
	fmt.Printf("Creating game for userId '%d'...\n", userId)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}
	createdGame := a.Game.CreateGame(userId)
	fmt.Println("Created game:", createdGame)
	c.JSON(http.StatusOK, createdGame)
}

func (a Api) UpdateGame(c *gin.Context) {
	gameId := getGameIdFromParam(c)
	fmt.Printf("Updating game '%d'...\n", gameId)
	c.Status(http.StatusOK)
}

func (a Api) DeleteGame(c *gin.Context) {
	gameId := getGameIdFromParam(c)
	fmt.Printf("Deleting game '%d'...\n", gameId)
	c.Status(http.StatusOK)
}

func (a Api) GuessWord(c *gin.Context) {
	fmt.Println("Guessing word...")
	gameId := getGameIdFromParam(c)
	userId, err := getUserIdFromCookie(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}

	var guess models.Guess
	c.BindJSON(&guess)
	fmt.Printf("Guessing word %q for game '%d'\n", guess.Word, gameId)
	err = a.ValidateGuess(guess.Word, gameId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	wordDetails, getWordErr := a.ValidateWordExists(guess.Word)
	if getWordErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": getWordErr.Error()})
		return
	}

	results := a.Game.GuessWord(gameId, guess.Word)
	gameover := a.Game.CheckGameOver(gameId)
	gameover.AnswerDescription = a.getAnswerDetails(gameover, wordDetails)

	dto := dto.GuessResponse{Word: guess.Word, Results: results, Gameover: gameover}
	c.JSON(http.StatusOK, dto)
}

func (a Api) getAnswerDetails(gameover models.Gameover, wordDetails wordsApi.WordDetails) string {
	isGameover := gameover.IsGameover
	answer := gameover.Answer
	isWon := gameover.Win

	if !isGameover {
		return ""
	}

	detailsUndefined := reflect.DeepEqual(wordDetails, wordsApi.WordDetails{})
	haveDetailsAlready := isWon && !detailsUndefined
	if haveDetailsAlready {
		fmt.Println("Already have details for answer.")
		return wordDetails.Results[0].Definition
	}

	answerDetails, err := a.WordsApi.GetWord(answer)
	if err != nil {
		fmt.Println("Error getting word:", err.Error(), ", using default word details")
		answerDetails = wordsApi.GetDefaultWordDetails(answer)
	}
	return answerDetails.Results[0].Definition
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
