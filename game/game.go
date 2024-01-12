package game

import "strings"

const length = 6

type Game struct {
}

func CheckWord(guess string, answer string) []string {
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
