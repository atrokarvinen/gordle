package api

import (
	"fmt"
	"gordle/dictionaryClient"
	"gordle/game/answers"
	"os"
	"sort"
	"strings"
)

func RemoveDuplicates(apiClient dictionaryClient.IDictionaryClient, lang string, wordLength int) {
	langCapitalized := strings.ToUpper(lang[:1]) + lang[1:]
	words := answers.GetAnswers(lang, "easy", wordLength)
	saveFile := fmt.Sprintf("./game/answers/answers_%s_%d.go", lang, wordLength)

	fmt.Println("Removing duplicates from", saveFile)

	uniqueWordsMap := make(map[string]bool)
	for _, word := range words {
		uniqueWordsMap[word] = true
	}
	uniqueWords := []string{}
	for word := range uniqueWordsMap {
		uniqueWords = append(uniqueWords, word)
	}
	sort.Strings(uniqueWords)

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
	_, err = file.WriteString(formatStringArray(uniqueWords))
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

	fmt.Println("Word count reduced from", len(words), "to", len(uniqueWords))
	fmt.Println("Removed duplicates from", saveFile)
}
