package api

import (
	"fmt"
	"go-test/game"
	"go-test/wordsApi"
	"strings"
)

func (a Api) ValidateGuess(word string, gameId int, userId int) error {
	game, err := a.Game.GetGame(gameId)
	if err != nil {
		return fmt.Errorf("Game for game id '%d' not found", gameId)
	}
	if game.UserId != userId {
		return fmt.Errorf("User '%d' has no access to this game", userId)
	}
	if len(word) != game.WordLength {
		return fmt.Errorf("Guess length (%d) does not match word length (%d)", len(word), game.WordLength)
	}
	gameover := a.Game.CheckGameOver(gameId)
	if gameover.IsGameover {
		return fmt.Errorf("Game is already over")
	}
	return nil
}

func (a Api) ValidateWordExists(word string) (wordsApi.WordDetails, error) {
	// Check if word is in the list of answers
	for _, answer := range game.Answers {
		if strings.ToLower(word) == strings.ToLower(answer) {
			fmt.Printf("Word %q is in the list of answers\n", word)
			return wordsApi.WordDetails{}, nil
		}
	}

	// Check word from the Words API
	wordDetails, err := a.WordsApi.GetWord(word)
	if err != nil {
		fmt.Println("Error getting word:", err)
	}
	if err != nil && err.Error() == fmt.Sprintf("Word '%s' not found", word) {
		return wordsApi.WordDetails{}, err
	}

	return wordDetails, nil
}
