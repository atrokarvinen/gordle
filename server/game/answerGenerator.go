package game

import (
	"fmt"
	"gordle/game/answers"
	"math/rand"
	"slices"
	"strings"
)

func GenerateRandomAnswer(lang string, wordLength int, playedAnswers []string) string {
	allAnswers := answers.GetAnswers(lang, wordLength)
	allAnswers = strListToLower(allAnswers)
	playedAnswers = strListToLower(playedAnswers)
	leastPlayedAnswers := GetLeastPlayedAnswers(playedAnswers, allAnswers)
	randomIndex := rand.Intn(len(leastPlayedAnswers))
	answer := leastPlayedAnswers[randomIndex]
	return strings.ToLower(answer)
}

func strListToLower(list []string) []string {
	for i, str := range list {
		list[i] = strings.ToLower(str)
	}
	return list
}

func GetLeastPlayedAnswers(playedAnswers []string, allAnswers []string) []string {
	playCountByAnswer := make(map[string]int)
	for _, answer := range allAnswers {
		playCountByAnswer[answer] = 0
	}
	for _, answer := range playedAnswers {
		playCountByAnswer[answer]++
	}
	minPlayCount := 9999
	for _, answer := range allAnswers {
		if playCountByAnswer[answer] < minPlayCount {
			minPlayCount = playCountByAnswer[answer]
		}
	}
	fmt.Println("minPlayCount", minPlayCount)
	fmt.Printf("Player has played %d / %d answers\n", len(playedAnswers), len(allAnswers))
	leastPlayedAnswers := []string{}
	for answer, playCount := range playCountByAnswer {
		if playCount == minPlayCount {
			leastPlayedAnswers = append(leastPlayedAnswers, answer)
		}
	}
	fmt.Printf("Player has %d least played answers left\n", len(leastPlayedAnswers))
	slices.Sort(leastPlayedAnswers)
	return leastPlayedAnswers
}
