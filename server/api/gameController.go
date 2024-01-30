package api

import (
	"fmt"
	"go-test/models"
	"go-test/models/dbModels"
	"go-test/models/dto"
	"go-test/wordsApi"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (a Api) GetGame(c *gin.Context) {
	userId, gameId, err := getIdsFromContext(c)
	if err != nil {
		return
	}
	fmt.Printf("Getting game '%d'...\n", gameId)
	game, err := a.Game.GetGame(gameId)
	err = a.ValidateGameFoundAndHasAccess(c, game, err, userId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, game)
}

func (a Api) GetLatestGame(c *gin.Context) {
	userId, err := getUserIdFromCookie(c)
	fmt.Println("Getting latest game for user '", userId, "'...")
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}
	game, err := a.Game.GetLatestGame(userId)
	err = a.ValidateGameFoundAndHasAccess(c, game, err, userId)
	if err != nil {
		return
	}
	gameover := a.Game.CheckGameOver(game.Id)
	dto := game
	dto.Gameover = gameover
	c.JSON(http.StatusOK, dto)
}

func (a Api) CreateGame(c *gin.Context) {
	userId, err := getUserIdFromCookie(c)
	fmt.Printf("Creating game for userId '%d'...\n", userId)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}
	request := dto.CreateGameRequest{}
	err = c.Bind(&request)
	if err != nil || reflect.DeepEqual(request, dto.CreateGameRequest{}) {
		fmt.Println("Error binding request, using default request")
		request = dto.DefaultCreateGameRequest
	}
	createdGame := a.Game.CreateGame(userId, request)
	fmt.Println("Created game:", createdGame)
	c.JSON(http.StatusOK, createdGame)
}

func (a Api) UpdateGame(c *gin.Context) {
	gameId := getGameIdFromParam(c)
	fmt.Printf("Updating game '%d'...\n", gameId)
	c.Status(http.StatusOK)
}

func (a Api) DeleteGame(c *gin.Context) {
	userId, gameId, err := getIdsFromContext(c)
	if err != nil {
		return
	}
	game, err := a.Game.GetGame(gameId)
	err = a.ValidateGameFoundAndHasAccess(c, game, err, userId)
	if err != nil {
		return
	}
	fmt.Printf("Deleting game '%d'...\n", gameId)

	gameover := models.Gameover{IsGameover: true, Win: false, Answer: game.Answer}
	answerDescription := a.getAnswerDetails(gameover, wordsApi.WordDetails{})
	fmt.Println("Answer description:", answerDescription)
	game.Gameover.IsGameover = true
	game.Gameover.AnswerDescription = answerDescription

	game.AnswerDescription = answerDescription
	game.State = int(dbModels.Lose)
	err = a.Game.UpdateGame(game)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, game)
}

func (a Api) GuessWord(c *gin.Context) {
	fmt.Println("Guessing word...")
	userId, gameId, err := getIdsFromContext(c)
	if err != nil {
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
	game, _ := a.Game.GetGame(gameId)
	options := dto.CreateGameRequest{Language: game.Language, WordLength: game.WordLength}
	wordDetails, getWordErr := a.ValidateWordExists(guess.Word, options)
	if getWordErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": getWordErr.Error()})
		return
	}

	results := a.Game.GuessWord(gameId, guess.Word)
	gameover := a.Game.CheckGameOver(gameId)
	gameover.AnswerDescription = a.getAnswerDetails(gameover, wordDetails)
	if gameover.IsGameover {
		game.AnswerDescription = gameover.AnswerDescription
		if gameover.Win {
			game.State = int(dbModels.Win)
		} else {
			game.State = int(dbModels.Lose)
		}
		a.Game.UpdateGame(game)
	}

	dto := dto.GuessResponse{Word: guess.Word, Results: results, Gameover: gameover}
	c.JSON(http.StatusOK, dto)
}

func (a Api) getAnswerDetails(gameover models.Gameover, wordDetails wordsApi.WordDetails) string {
	isGameover := gameover.IsGameover
	answer := gameover.Answer
	isWon := gameover.Win

	if !isGameover {
		fmt.Println("Game is not over, not getting answer details")
		return ""
	}

	detailsUndefined := reflect.DeepEqual(wordDetails, wordsApi.WordDetails{})
	haveDetailsAlready := isWon && !detailsUndefined
	if haveDetailsAlready {
		fmt.Println("Already have details for answer.")
		return parseDefinition(wordDetails.Results)
	}

	answerDetails, err := a.WordsApi.GetWord(answer)
	if err != nil {
		fmt.Println("Error getting word:", err.Error(), ", using default word details")
		answerDetails = wordsApi.GetDefaultWordDetails(answer)
	}
	return parseDefinition(answerDetails.Results)
}

func parseDefinition(results []wordsApi.WordResults) string {
	if len(results) == 0 {
		return "Definition not found"
	}
	definitions := []string{}
	for _, result := range results {
		definitions = append(definitions, result.Definition)
	}
	return strings.Join(definitions, ";;")
}

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
