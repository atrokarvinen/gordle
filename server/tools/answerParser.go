package tools

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode/utf8"
)

// https://www.kotus.fi/aineistot/sana-aineistot/nykysuomen_sanalista

func ParseCsvAnswers() {
	records, err := ReadCsvRecords("nykysuomensanalista2022.csv")
	if err != nil {
		fmt.Println("Error reading records: " + err.Error())
		return
	}
	wordLengths := []int{5, 6, 7, 8}
	for _, length := range wordLengths {
		SaveAnswers(records, length, "fi")
	}
}

func ReadCsvRecords(fileName string) ([][]string, error) {
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

func ParseTxtAnswers() {
	isOneLine := true
	isEasy := true

	// file := "answers_swe_easy.txt"
	// lang := "se"

	// file := "answers_german_easy.txt"
	// lang := "de"

	// file := "dictionary-pl.txt"
	// lang := "pl"

	file := "answers_fi_easy2.txt"
	lang := "fi"

	records, err := ReadTxtRecords(file, isOneLine)
	if err != nil {
		fmt.Println("Error reading records: " + err.Error())
		return
	}
	n := 2000
	top_n_records := GetTopNRecords(records, n)
	wordLengths := []int{5, 6, 7, 8}
	for _, length := range wordLengths {
		SaveAnswers(top_n_records, length, lang, isEasy)
	}
}

func ReadTxtRecords(fileName string, isOneLine bool) ([][]string, error) {
	fmt.Println("Parsing answers from file '" + fileName + "'...")

	data, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error reading file: " + err.Error())
		return nil, err
	}

	scanner := bufio.NewScanner(data)
	var records [][]string
	if isOneLine {
		scanner.Split(bufio.ScanWords)
	}
	for scanner.Scan() {
		line := scanner.Text()
		records = append(records, []string{line})
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading txt file: " + err.Error())
		return nil, err
	}
	fmt.Println("Parsed", len(records), "records")
	return records, nil
}

func SaveAnswers(records [][]string, length int, lang string, isEasy ...bool) {
	langCapitalized := strings.ToUpper(lang[0:1]) + lang[1:]
	easyString := ""
	easyStringCapitalized := ""
	if len(isEasy) > 0 && isEasy[0] {
		easyString = "_easy"
		easyStringCapitalized = "Easy"
	}
	fileName := fmt.Sprintf("answers_%s%s_%d.go", lang, easyString, length)
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file: " + err.Error())
		return
	}

	defer f.Close()

	words := GetWordsOfLength(records, length, lang)
	words = GetUniqueWords(words)

	fmt.Fprintln(f, "package answers")
	fmt.Fprintf(f, "\nvar Answers%s%s%d = []string{\n", langCapitalized, easyStringCapitalized, length)
	for _, word := range words {
		text := fmt.Sprintf("\t\"%s\",\n", word)
		fmt.Print(text)
		f.WriteString(text)
	}
	f.WriteString("}")

	fmt.Printf("Saved %d words of length %d to file '%s'\n", len(words), length, fileName)
}

func GetWordsOfLength(records [][]string, length int, lang string) []string {
	words := []string{}
	for _, record := range records {
		// fmt.Println("Record:", record)
		word := record[0]
		correctLength := utf8.RuneCountInString(word) == length
		if !correctLength {
			continue
		}
		alphabets := GetAlphabets(lang)
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

func GetAlphabets(lang string) string {
	switch lang {
	case "se":
		return "abcdefghijklmnopqrstuvwxyzåäö"
	case "fi":
		return "abcdefghijklmnopqrstuvwxyzåäö"
	case "de":
		return "abcdefghijklmnopqrstuvwxyzäöüß"
	case "en":
		return "abcdefghijklmnopqrstuvwxyz"
	case "pl":
		return "aąbcćdeęfghijklłmnńoóprsśtuwyzźż"
	}
	panic(fmt.Sprintf("Unknown language '%s'", lang))
}

func GetTopNRecords(records [][]string, n int) [][]string {
	if len(records) < n {
		return records
	}
	return records[:n]
}

func GetUniqueWords(words []string) []string {
	uniqueWordsMap := make(map[string]bool)
	for _, word := range words {
		uniqueWordsMap[word] = true
	}
	uniqueWords := []string{}
	for word := range uniqueWordsMap {
		uniqueWords = append(uniqueWords, word)
	}
	sort.Strings(uniqueWords)
	return uniqueWords
}
