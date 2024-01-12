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

	for {
		previousGuesses := newGame.GetGuesses()
		fmt.Println("\nPrevious guesses:\n", strings.Join(previousGuesses, "\n"))

		guess := RequestGuessFromUser()

		fmt.Println("Guess:", guess)

		results := newGame.CheckWord(guess, answer)

		fmt.Println(strings.Join(strings.Split(answer, ""), ","))
		fmt.Println(strings.Join(strings.Split(guess, ""), ","))
		fmt.Println(strings.Join(results, ","))

		newGame.DataProvider.AddGuess(guess)

		isCorrect := game.IsCorrectResult(results)
		if isCorrect {
			fmt.Println("Correct!")
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
