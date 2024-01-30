package tools

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// https://www.kotus.fi/aineistot/sana-aineistot/nykysuomen_sanalista

func ParseAnswers() {
	records, err := ReadRecords("nykysuomensanalista2022.csv")
	if err != nil {
		fmt.Println("Error reading records: " + err.Error())
		return
	}
	wordLengths := []int{5, 6, 7, 8}
	for _, length := range wordLengths {
		SaveAnswers(records, length)
	}
}

func ReadRecords(fileName string) ([][]string, error) {
	fmt.Println("Parsing answers from file '" + fileName + "'...")

	data, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error reading file: " + err.Error())
		return nil, err
	}

	csvReader := csv.NewReader(data)
	csvReader.Comma = '\t'
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading csv: " + err.Error())
		return nil, err
	}
	return records, nil
}

func SaveAnswers(records [][]string, length int) {
	fileName := fmt.Sprintf("answers_fi_%d.go", length)
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file: " + err.Error())
		return
	}

	defer f.Close()

	words := GetWordsOfLength(records, length)

	f.WriteString(fmt.Sprintln("package answers"))
	f.WriteString(fmt.Sprintf("var AnswersFi%d = []string{\n", length))
	for _, word := range words {
		text := fmt.Sprintf("\"%s\",\n", word)
		fmt.Print(text)
		f.WriteString(text)
	}
	f.WriteString("}")
}

func GetWordsOfLength(records [][]string, length int) []string {
	words := []string{}
	for _, record := range records {
		word := record[0]
		correctLength := utf8.RuneCountInString(word) == length
		if !correctLength {
			continue
		}
		alphabets := "abcdefghijklmnopqrstuvwxyzåäö"
		wordChars := strings.Split(word, "")
		validCharacters := true
		for _, char := range wordChars {
			if !strings.ContainsAny(alphabets, char) {
				validCharacters = false
			}
		}
		if !validCharacters {
			continue
		}
		words = append(words, record[0])
	}
	return words
}
