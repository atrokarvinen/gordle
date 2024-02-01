package api

import (
	"fmt"
	"go-test/models"
	"go-test/models/dbModels"
	"go-test/models/dto"
	"go-test/wordsApi"
	"net/http"
	"reflect"
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
	answerDescription := a.getAnswerDetails(gameover, wordsApi.DictionaryDetails{}, game.Language)
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

func (a Api) getAnswerDetails(gameover models.Gameover, wordDetails wordsApi.DictionaryDetails, lang string) string {
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
		return strings.Join(wordDetails.Definitions, ";;")
	}

	answerDetails, err := a.DictionaryFactory.GetDictionaryClient(lang).GetWord(answer)
	if err != nil {
		fmt.Println("Error getting word:", err.Error(), ", using default word details")
		answerDetails = wordsApi.GetDefaultWordDetails(answer)
	}
	return strings.Join(answerDetails.Definitions, ";;")
}

// <span class="kt-scope">halv.</span> laiha, huono hevonen, koni, luuska, kopukka.
