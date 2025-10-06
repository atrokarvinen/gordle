package api

import (
	"fmt"
	"gordle/models"
	"gordle/models/dbModels"
	"gordle/models/dto"
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
		errorDto := ApiError{
			Message: getWordErr.Error(),
			Data:    map[string]string{"word": guess.Word, "language": game.Language}}
		c.JSON(http.StatusBadRequest, errorDto)
		return
	}

	results := a.Game.GuessWord(gameId, guess.Word)
	gameover := a.Game.CheckGameOver(gameId)
	answerDetails := a.getAnswerDetails(gameover, wordDetails, game.Language, game.Difficulty)
	gameover.Definitions = answerDetails.Definitions
	gameover.Examples = answerDetails.Examples
	if gameover.IsGameover {
		game.AnswerDescription = gameover.Definitions
		game.AnswerExamples = gameover.Examples
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
