package api

import (
	"fmt"
	"gordle/dictionaryClient"
	"gordle/game/answers"
	"os"
	"strings"
	"time"
)

func VerifyWords(apiClient dictionaryClient.IDictionaryClient, lang string, wordLength int) {
	langCapitalized := strings.ToUpper(lang[:1]) + lang[1:]
	words := answers.GetAnswers(lang, "easy", wordLength)
	saveFile := fmt.Sprintf("./game/answers/answers_%s_%d.go", lang, wordLength)
	top10Words := words[0:10]
	words = top10Words

	valids := []string{}
	invalids := []string{}
	wordIndex := 0
	for _, word := range words {
		wordIndex++
		fmt.Printf("Checking word %d / %d: '%s'\n", wordIndex, len(words), word)

		word := strings.Trim(word, " ")
		details, err := apiClient.GetWord(word, lang)
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

		hasDefinitions := len(details.Definitions) > 0
		hasExamples := len(details.Examples) > 0
		if !hasDefinitions && !hasExamples {
			fmt.Printf("Word '%s' is invalid, no definitions or examples\n", word)
			invalids = append(invalids, word)
			continue
		}

		fmt.Printf("Word '%s' is valid\n", word)
		valids = append(valids, word)
	}

	// formattedValids := formatStringArray(valids)
	file, err := os.OpenFile(saveFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	err = file.Truncate(0)
	if err != nil {
		fmt.Println("Error truncating file:", err)
		return
	}
	_, err = file.WriteString("package answers\n\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	_, err = file.WriteString(fmt.Sprintf("var Answers%s%d = []string{\n", langCapitalized, wordLength))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	_, err = file.WriteString(formatStringArray(valids))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	_, err = file.WriteString("}\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	err = file.Close()
	if err != nil {
		fmt.Println("Error closing file:", err)
		return
	}

	formattedInvalids := formatStringArray(invalids)
	err = os.WriteFile("invalids.txt", []byte(fmt.Sprintf("%v", formattedInvalids)), 0644)
	if err != nil {
		fmt.Println("Error writing invalids.txt", err)
	}

	fmt.Printf("Valid words: %d / %d\n", len(valids), len(words))
	fmt.Printf("Invalid words: %d / %d\n", len(invalids), len(words))
	fmt.Println("Done")
}

func formatStringArray(sArr []string) string {
	var formatted string = ""
	for _, word := range sArr {
		formatted += "\t\"" + word + "\"," + "\n"
	}
	return formatted
}
