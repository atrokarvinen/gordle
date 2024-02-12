package api

import (
	"fmt"
	"gordle/dictionaryClient"
	"gordle/game"
	"os"
	"strings"
	"time"
)

func VerifyWords(apiClient dictionaryClient.IDictionaryClient) {
	words := game.RandomWords

	valids := []string{}
	invalids := []string{}
	for _, word := range words {
		word := strings.Trim(word, " ")
		_, err := apiClient.GetWord(word)
		if err != nil && err.Error() == "Too many requests" {
			fmt.Println("Too many requests")
			break
		}
		if err != nil && err.Error() == "invalid character 'T' looking for beginning of value" {
			fmt.Println("Too many requests, sleeping for 4 minutes...")
			time.Sleep(4 * time.Minute)
			continue
		}
		if err != nil {
			fmt.Println("Error getting word", err)
			fmt.Printf("Word '%s' is invalid\n", word)
			invalids = append(invalids, word)
			continue
		}
		fmt.Printf("Word '%s' is valid\n", word)
		valids = append(valids, word)
	}

	formattedValids := formatStringArray(valids)
	err := os.WriteFile("valids.txt", []byte(fmt.Sprintf("%v", formattedValids)), 0644)
	if err != nil {
		fmt.Println("Error writing valids.txt", err)
	}

	formattedInvalids := formatStringArray(invalids)
	err = os.WriteFile("invalids.txt", []byte(fmt.Sprintf("%v", formattedInvalids)), 0644)
	if err != nil {
		fmt.Println("Error writing invalids.txt", err)
	}
}

func formatStringArray(sArr []string) string {
	var formatted string = ""
	for _, word := range sArr {
		formatted += "\"" + word + "\"," + "\n"
	}
	return formatted
}
