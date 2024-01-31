package game

import (
	"fmt"
	"go-test/game/answers"
	"go-test/models"
	"go-test/models/dbModels"
	"go-test/models/dto"
	"go-test/wordsApi"
	"math/rand"
	"strings"
	"unicode/utf8"
)

type Game struct {
	DataProvider DataProvider
	WordsApi     wordsApi.IWordsApiClient
}

func (g Game) CreateGame(userId int, gameOptions dto.CreateGameRequest) models.Game {
	answer := g.GenerateRandomAnswer(gameOptions.Language, gameOptions.WordLength)
	fmt.Printf("Creating game: Answer=%s, MaxAttempts=%d, WordLength=%d, Language=%s...\n",
		answer, gameOptions.MaxAttempts, gameOptions.WordLength, gameOptions.Language)

	game := dbModels.Game{
		Answer:      answer,
		MaxAttempts: gameOptions.MaxAttempts,
		WordLength:  gameOptions.WordLength,
		Language:    gameOptions.Language,
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

func (g Game) GenerateRandomAnswer(lang string, wordLength int) string {
	allAnswers := answers.GetAnswers(lang, wordLength)
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

func (g Game) CheckWord(guessStr string, answerStr string) []string {
	length := utf8.RuneCountInString(answerStr)
	results := make([]string, length)
	guess := []rune(strings.ToLower(guessStr))
	answer := []rune(strings.ToLower(answerStr))

	fmt.Println("Checking word", guessStr, "against", answerStr)

	// Find correct letters
	var containedMap = make(map[rune]int)
	for i := 0; i < length; i++ {
		answerChar := answer[i]
		guessChar := guess[i]
		isGuessCorrect := answerChar == guessChar
		if isGuessCorrect {
			containedMap[guessChar]++
			results[i] = "v"
		}
	}

	// Find contained letters
	for i := 0; i < length; i++ {
		guessChar := guess[i]
		if results[i] == "v" {
			continue
		}
		count := 0
		for j := 0; j < length; j++ {
			answerChar := answer[j]
			if guessChar == answerChar {
				count++
			}
		}
		if count > 0 {
			if containedMap[guessChar] >= count {
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
