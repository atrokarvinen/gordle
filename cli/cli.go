package cli

import (
	"bufio"
	"fmt"
	"go-test/game"
	"go-test/models"
	"os"
	"strings"
)

type Cli struct {
}

func (c Cli) Run() {
	settings := game.GameSettings{MaxAttempts: 6, WordLength: 6}
	// dataProvider := game.InMemoryDataProvider{}
	dataProvider := game.DatabaseDataProvider{}
	newGame := game.Game{DataProvider: dataProvider, Settings: settings}

	gameId := 1

	answer := newGame.GenerateRandomAnswer()
	for {
		guessedWord := RequestGuessFromUser(newGame.Settings.WordLength)
		guess := models.Guess{Word: guessedWord, GameId: gameId}

		// fmt.Println("Guess:", guess)

		results := newGame.CheckWord(guessedWord, answer)

		// fmt.Println(strings.Join(strings.Split(answer, ""), ","))
		// fmt.Println(strings.Join(strings.Split(guess, ""), ","))
		// fmt.Println(strings.Join(results, ","))

		newGame.DataProvider.AddGuess(guess)
		previousGuesses := newGame.GetGuesses(gameId)
		guessedWords := make([]string, len(previousGuesses))
		for i, guess := range previousGuesses {
			guessedWords[i] = guess.Word
		}
		PrintGameState(guessedWords, answer, newGame)

		isCorrect := game.IsCorrectResult(results)
		guessCount := len(previousGuesses)

		if isCorrect {
			fmt.Println("Correct!")
			break
		} else if guessCount >= newGame.Settings.MaxAttempts {
			fmt.Println("You lose! Answer was:", answer)
			break
		}
	}
}

func RequestGuessFromUser(length int) string {
	for {
		fmt.Printf("\nEnter guessed word of length (%d):\n", length)
		text := Read()
		if len(text) != 6 {
			fmt.Printf("Word %q has length (%d). Expected (%d)\n", text, len(text), length)
			continue
		}
		return text
	}
}

func Read() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	trimmedText := strings.Trim(text, "\r\n")
	return trimmedText
}

func PrintGameState(guesses []string, answer string, newGame game.Game) {
	fmt.Println("\nGame state:")

	for i, guess := range guesses {
		results := newGame.CheckWord(guess, answer)
		fmt.Println(i+1, "/ 6 \t", guess, "\t => \t", strings.Join(results, ","))
	}

	fmt.Println()
}
