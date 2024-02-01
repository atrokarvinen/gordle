package api

import (
	"fmt"
	"go-test/models"
	"go-test/models/dbModels"
	"go-test/models/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	gameover.AnswerDescription = a.getAnswerDetails(gameover, wordDetails, game.Language)
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
