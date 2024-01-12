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
	for i := 0; ; i++ {
		var guess string
		for i := 0; ; i++ {
			text := Read()
			if len(text) != 6 {
				fmt.Printf("Word %q has length (%d). Expected (%d)\n", text, len(text), length)
				continue
			}
			guess = text
			break
		}

		// var game game.Game = game.Game{}

		fmt.Println("Guess:", guess)

		answer := "gordle"

		results := game.CheckWord(guess, answer)

		fmt.Println(strings.Join(strings.Split(answer, ""), ","))
		fmt.Println(strings.Join(strings.Split(guess, ""), ","))
		fmt.Println(strings.Join(results, ","))

		isCorrect := game.IsCorrectResult(results)
		if isCorrect {
			fmt.Println("Correct!")
			break
		}
	}
}

func Read() string {
	fmt.Println("\nEnter guessed word:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	trimmedText := strings.Trim(text, "\r\n")
	return trimmedText
}
