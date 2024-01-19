package game

import (
	"fmt"
	"go-test/models"
	"math/rand"
	"strings"
)

type Game struct {
	DataProvider DataProvider
}

func (g Game) CreateGame(name string) models.GameType {
	answer := g.GenerateRandomAnswer()
	fmt.Println("answer", answer)
	game := models.GameType{Answer: answer, Name: name, MaxAttempts: 6, WordLength: 6}
	createdGame := g.DataProvider.CreateGame(game)
	return createdGame
}

func (g Game) GetLatestGame() models.GameType {
	return g.DataProvider.GetLatestGame()
}
func (g Game) GetGames() []models.GameType {
	return g.DataProvider.GetGames()
}

func (g Game) LoadGame(gameId int) (models.GameType, error) {
	fmt.Printf("Loading game '%d'...\n", gameId)
	return g.DataProvider.GetGame(gameId)
}

func (g Game) GetGuesses(gameId int) []models.Guess {
	return g.DataProvider.GetPreviousGuesses(gameId)
}

func (g Game) GenerateRandomAnswer() string {
	var allAnswers = answers
	randomIndex := rand.Intn(len(allAnswers))
	answer := allAnswers[randomIndex]
	return strings.ToLower(answer)
}

func (g Game) GuessWord(gameId int, guess string) []string {
	game, _ := g.DataProvider.GetGame(gameId)
	results := g.CheckWord(guess, game.Answer, game.WordLength)
	g.DataProvider.AddGuess(models.Guess{Word: guess, GameId: gameId})
	return results
}

func (g Game) CheckGameOver(gameId int) models.Gameover {
	game, _ := g.DataProvider.GetGame(gameId)
	guesses := g.DataProvider.GetPreviousGuesses(gameId)
	isGameWon := IsCorrectResult(g.CheckWord(guesses[len(guesses)-1].Word, game.Answer, game.WordLength))
	isGameOver := len(guesses) >= game.MaxAttempts || isGameWon
	fmt.Println("Is game over?", isGameOver, "Is game won?", isGameWon, "attempts", len(guesses), "max attempts", game.MaxAttempts)
	return models.Gameover{IsGameover: isGameOver, Win: isGameWon, Answer: game.Answer, AnswerDescription: "description"}
}

func (g Game) CheckWord(guess string, answer string, length int) []string {
	results := make([]string, length)

	// Find correct letters
	var containedMap = make(map[string]int)
	for i := 0; i < length; i++ {
		answerChar := strings.ToLower(string(answer[i]))
		guessChar := strings.ToLower(string(guess[i]))
		isGuessCorrect := answerChar == guessChar
		if isGuessCorrect {
			containedMap[guessChar]++
			results[i] = "v"
		}
	}

	// Find contained letters
	for i := 0; i < length; i++ {
		guessChar := strings.ToLower(string(guess[i]))
		if results[i] == "v" {
			continue
		}
		isGuessContained := strings.Contains(answer, guessChar)
		if isGuessContained {
			if containedMap[guessChar] >= strings.Count(answer, guessChar) {
				results[i] = "x"
			} else {
				containedMap[guessChar]++
				results[i] = "?"
			}
		}
	}

	// Fill with incorrect letters
	for i := 0; i < length; i++ {
		if results[i] == "" {
			results[i] = "x"
		}
	}

	return results[:]
}

func IsCorrectResult(results []string) bool {
	for _, result := range results {
		if result != "v" {
			return false
		}
	}
	return true
}

func (g Game) IsGameOver(gameId int) bool {
	guesses := g.DataProvider.GetPreviousGuesses(gameId)
	game, _ := g.DataProvider.GetGame(gameId)
	return len(guesses) >= game.MaxAttempts
}
