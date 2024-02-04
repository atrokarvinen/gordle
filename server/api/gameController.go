package api

import (
	"fmt"
	m "gordle/dictionaryClient/models"
	"gordle/models"
	"gordle/models/dbModels"
	"gordle/models/dto"
	"net/http"
	"reflect"

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
	userId, err := getUserIdFromContext(c)
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
	userId, err := getUserIdFromContext(c)
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
	answerDetails := a.getAnswerDetails(gameover, m.DictionaryDetails{}, game.Language)
	fmt.Println("Answer description:", answerDetails)
	game.Gameover.IsGameover = true
	game.Gameover.Definitions = answerDetails.Definitions
	game.Gameover.Examples = answerDetails.Examples

	game.AnswerDescription = answerDetails.Definitions
	game.AnswerExamples = answerDetails.Examples
	game.State = int(dbModels.Lose)
	err = a.Game.UpdateGame(game)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, game)
}

func (a Api) getAnswerDetails(gameover models.Gameover, wordDetails m.DictionaryDetails, lang string) m.DictionaryDetails {
	isGameover := gameover.IsGameover
	answer := gameover.Answer
	isWon := gameover.Win

	if !isGameover {
		fmt.Println("Game is not over, not getting answer details")
		return m.DictionaryDetails{}
	}

	detailsUndefined := reflect.DeepEqual(wordDetails, m.DictionaryDetails{})
	haveDetailsAlready := isWon && !detailsUndefined
	if haveDetailsAlready {
		fmt.Println("Already have details for answer.")
		return wordDetails
	}

	answerDetails, err := a.DictionaryFactory.GetDictionaryClient(lang).GetWord(answer)
	if err != nil {
		fmt.Println("Error getting word:", err.Error(), ", using default word details")
		answerDetails = m.GetDefaultWordDetails(answer)
	}
	return answerDetails
}
