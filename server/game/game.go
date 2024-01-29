package game

import (
	"fmt"
	"go-test/models"
	"go-test/models/dbModels"
	"go-test/wordsApi"
	"math/rand"
	"strings"
)

type Game struct {
	DataProvider DataProvider
	WordsApi     wordsApi.IWordsApiClient
}

func (g Game) CreateGame(userId int) models.Game {
	answer := g.GenerateRandomAnswer()
	fmt.Println("answer", answer)

	game := dbModels.Game{
		Answer:      answer,
		MaxAttempts: 6,
		WordLength:  6,
		UserID:      userId,
		State:       int(dbModels.Active),
	}

	createdGame := g.DataProvider.CreateGame(game)
	dto := g.MapDbGameToGame(createdGame)
	return dto
}

func (g Game) GetLatestGame(userId int) (models.Game, error) {
	game, err := g.DataProvider.GetLatestGame(userId)
	if err != nil {
		return models.Game{}, err
	}
	return g.MapDbGameToGame(game), nil
}

func (g Game) GetGame(gameId int) (models.Game, error) {
	fmt.Printf("Loading game '%d'...\n", gameId)
	game, err := g.DataProvider.GetGame(gameId)
	if err != nil {
		return models.Game{}, err
	}
	return g.MapDbGameToGame(game), nil
}

func (g Game) GenerateRandomAnswer() string {
	var allAnswers = Answers
	randomIndex := rand.Intn(len(allAnswers))
	answer := allAnswers[randomIndex]
	return strings.ToLower(answer)
}

func (g Game) GuessWord(gameId int, guess string) []string {
	game, _ := g.DataProvider.GetGame(gameId)
	results := g.CheckWord(guess, game.Answer)
	guessToCreate := dbModels.Guess{Word: guess, GameID: gameId}
	g.DataProvider.AddGuess(guessToCreate)
	return results
}

func (g Game) CheckGameOver(gameId int) models.Gameover {
	game, _ := g.DataProvider.GetGame(gameId)
	guesses := g.DataProvider.GetPreviousGuesses(gameId)
	isGameWon := false
	if len(guesses) > 0 {
		isGameWon = IsCorrectResult(g.CheckWord(guesses[len(guesses)-1].Word, game.Answer))
	}
	isGameOver := len(guesses) >= game.MaxAttempts || isGameWon
	fmt.Println("Is game over?", isGameOver, "Is game won?", isGameWon, "attempts", len(guesses), "max attempts", game.MaxAttempts)
	var answer string = ""
	if isGameOver {
		answer = game.Answer
	}
	return models.Gameover{IsGameover: isGameOver, Win: isGameWon, Answer: answer, AnswerDescription: ""}
}

func (g Game) UpdateGame(game models.Game) error {
	fmt.Printf("Updating game '%d', state '%d'.", game.Id, game.State)
	gameDb := g.MapGameToDbGame(game)
	return g.DataProvider.UpdateGame(gameDb)
}

func (g Game) CheckWord(guess string, answer string) []string {
	length := len(answer)
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
