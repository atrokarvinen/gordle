package game

import (
	"math/rand"
	"strings"
)

const length = 6

type Game struct {
	DataProvider DataProvider
}

func (g Game) GetGuesses() []string {
	return g.DataProvider.GetPreviousGuesses()
}

func (g Game) GenerateRandomAnswer() string {
	var allAnswers = answers
	randomIndex := rand.Intn(len(allAnswers))
	answer := allAnswers[randomIndex]
	return strings.ToLower(answer)
}

func (g Game) CheckWord(guess string, answer string) []string {
	var results [length]string

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
