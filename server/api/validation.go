package api

import (
	"fmt"
	m "gordle/dictionaryClient/models"
	"gordle/game/answers"
	"gordle/models"
	"gordle/models/dto"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

func (a Api) ValidateGuess(word string, gameId int, userId int) error {
	if gameId == -1 {
		return fmt.Errorf("No previous games found, please start a new one")
	}
	game, err := a.Game.GetGame(gameId)
	if err != nil {
		return fmt.Errorf("Game for game id '%d' not found", gameId)
	}
	if game.UserId != userId {
		return fmt.Errorf("User '%d' has no access to this game", userId)
	}
	if utf8.RuneCountInString(word) != game.WordLength {
		return fmt.Errorf("Guess length (%d) does not match word length (%d)", utf8.RuneCountInString(word), game.WordLength)
	}
	gameover := a.Game.CheckGameOver(gameId)
	if gameover.IsGameover {
		return fmt.Errorf("Game is already over")
	}
	return nil
}

func (a Api) ValidateWordExists(word string, gameOptions dto.CreateGameRequest) (m.DictionaryDetails, error) {
	// Check if word is in the list of answers
	difficulty := "all_words"
	wordLength := gameOptions.WordLength
	lang := gameOptions.Language
	allAnswers := answers.GetAnswers(lang, difficulty, wordLength)
	for _, answer := range allAnswers {
		if strings.EqualFold(word, answer) {
			fmt.Printf("Word %q is in the list of answers\n", word)
			return m.DictionaryDetails{}, nil
		}
	}

	// Check word from the dictionary API
	wordDetails, err := a.DictionaryFactory.GetDictionaryClient(lang, difficulty).GetWord(word, lang)
	if err != nil {
		fmt.Println("Error getting word:", err)
		return m.DictionaryDetails{}, fmt.Errorf("word_not_found")
	}

	hasDefitions := len(wordDetails.Definitions) > 0
	hasExamples := len(wordDetails.Examples) > 0
	if !hasDefitions && !hasExamples {
		fmt.Printf("Word %q has no definitions or examples\n", word)
		return m.DictionaryDetails{}, fmt.Errorf("word_not_found")
	}

	return wordDetails, nil
}

func (a Api) ValidateGameFoundAndHasAccess(c *gin.Context, game models.Game, getGameErr error, userId int) error {
	if getGameErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "game not found"})
		return getGameErr
	}
	if game.UserId != userId {
		c.JSON(http.StatusForbidden, gin.H{"message": "User has no access to this game"})
		return fmt.Errorf("User has no access to this game")
	}
	return nil
}
