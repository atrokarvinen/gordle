package api

import (
	"fmt"
	"go-test/dictionaryClient"
	"go-test/game"
	"os"
	"time"
)

func VerifyWords(apiClient dictionaryClient.IDictionaryClient) {
	words := game.RandomWords
	// words := []string{}
	// for i, word := range game.RandomWords {
	// 	if i > 5 {
	// 		break
	// 	}
	// 	words = append(words, word)
	// }

	valids := []string{}
	invalids := []string{}
	for _, word := range words {
		_, err := apiClient.GetWord(word)
		time.Sleep(100)
		if err != nil && err.Error() == "Too many requests" {
			fmt.Println("Too many requests")
			break
		}
		if err != nil {
			fmt.Println("Error getting word", err)
			fmt.Printf("Word %q is invalid\n", word)
			invalids = append(invalids, word)
			continue
		}
		fmt.Printf("Word %q is valid\n", word)
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
