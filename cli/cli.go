package cli

import (
	"bufio"
	"fmt"
	"go-test/game"
	"os"
	"strings"
)

type Cli struct {
}

const length = 6

func (c Cli) Run() {
	dataProvider := game.InMemoryDataProvider{}
	newGame := game.Game{DataProvider: dataProvider}

	answer := newGame.GenerateRandomAnswer()
	maxAttempts := 6

	for {
		guess := RequestGuessFromUser()

		// fmt.Println("Guess:", guess)

		results := newGame.CheckWord(guess, answer)

		// fmt.Println(strings.Join(strings.Split(answer, ""), ","))
		// fmt.Println(strings.Join(strings.Split(guess, ""), ","))
		// fmt.Println(strings.Join(results, ","))

		newGame.DataProvider.AddGuess(guess)
		previousGuesses := newGame.GetGuesses()
		PrintGameState(previousGuesses, answer, newGame)

		isCorrect := game.IsCorrectResult(results)
		guessCount := len(previousGuesses)

		if isCorrect {
			fmt.Println("Correct!")
			break
		} else if guessCount >= maxAttempts {
			fmt.Println("You lose! Answer was:", answer)
			break
		}
	}
}

func RequestGuessFromUser() string {
	for {
		text := Read()
		if len(text) != 6 {
			fmt.Printf("Word %q has length (%d). Expected (%d)\n", text, len(text), length)
			continue
		}
		return text
	}
}

func Read() string {
	fmt.Println("\nEnter guessed word:")
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
