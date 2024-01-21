package cli

import (
	"bufio"
	"fmt"
	"go-test/game"
	"go-test/models"
	"os"
	"strconv"
	"strings"
)

type Cli struct {
	DataProvider game.DataProvider
	Game         game.Game
}

func (c Cli) Run() {
	c.LobbyLoop()
}

func (c Cli) LobbyLoop() {
	for {
		fmt.Println("Welcome to the game!")
		fmt.Println("1. Start new game")
		fmt.Println("2. Continue game")
		fmt.Println("3. Exit")

		text := Read()

		if text == "1" {
			fmt.Println("Starting new game...")
			game := c.CreateGameLoop()
			c.GameLoop(game)
		} else if text == "2" {
			fmt.Println("Continuing game...")
			gameId := c.SelectGameLoop()
			game, err := c.Game.GetGame(gameId)
			if err != nil {
				fmt.Printf("Game '%d' not found\n", gameId)
				continue
			}
			c.PrintGame(game)
			c.GameLoop(game)
		} else if text == "3" {
			fmt.Println("Bye!")
			break
		} else {
			fmt.Println("Invalid option")
		}
	}
}

func (c Cli) CreateGameLoop() models.GameType {
	fmt.Println("Enter game name:")
	name := Read()
	game := c.Game.CreateGame(name)
	return game
}

func (c Cli) SelectGameLoop() int {
	for {
		fmt.Println("Select game:")

		games := c.Game.GetGames()
		for _, game := range games {
			fmt.Println(game.Id, game.Name)
		}

		text := Read()
		gameId, err := strconv.Atoi(text)

		if err != nil {
			fmt.Println("Invalid option: Input must be an integer")
			continue
		}

		return gameId
	}
}

func (c Cli) PrintGame(game models.GameType) {
	fmt.Println("Loaded game:")
	guesses := c.Game.GetGuesses(game.Id)
	guessedWords := make([]string, len(guesses))
	for i, guess := range guesses {
		guessedWords[i] = guess.Word
	}
	answer := game.Answer
	results := [][]string{}
	for _, guess := range guessedWords {
		result := c.Game.CheckWord(guess, answer, game.WordLength)
		results = append(results, result)
	}
	PrintGameState(guessedWords, answer, results)
}

func (c Cli) GameLoop(gameType models.GameType) {
	newGame := c.Game
	gameId := gameType.Id
	answer := gameType.Answer

	fmt.Printf("In game %q...\n", gameType.Name)

	for {
		guessedWord := RequestGuessFromUser(gameType.WordLength)
		if strings.ToLower(guessedWord) == "quit" {
			fmt.Println("Quitting game...")
			break
		}
		guess := models.Guess{Word: guessedWord, GameId: gameId}

		// fmt.Println("Guess:", guess)

		results := newGame.CheckWord(guessedWord, answer, gameType.WordLength)

		// fmt.Println(strings.Join(strings.Split(answer, ""), ","))
		// fmt.Println(strings.Join(strings.Split(guess, ""), ","))
		// fmt.Println(strings.Join(results, ","))

		newGame.DataProvider.AddGuess(guess)

		c.PrintGame(gameType)

		isCorrect := game.IsCorrectResult(results)
		isGameOver := c.Game.IsGameOver(gameId)

		if isCorrect {
			fmt.Println("Correct!")
			break
		} else if isGameOver {
			fmt.Println("You lose! Answer was:", answer)
			break
		}
	}
}

func RequestGuessFromUser(length int) string {
	for {
		fmt.Printf("To quit, enter \"quit\".\n")
		fmt.Printf("Enter guessed word of length (%d):\n", length)
		text := Read()
		isQuit := strings.ToLower(text) == "quit"
		if len(text) != 6 && !isQuit {
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

func PrintGameState(guesses []string, answer string, results [][]string) {
	fmt.Println("\nGame state:")

	for i, guess := range guesses {
		result := results[i]
		fmt.Println(i+1, "/ 6 \t", guess, "\t => \t", strings.Join(result, ","))
	}

	fmt.Println()
}
