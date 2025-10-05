package tools

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
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
	// file := "swe_wordlist.txt"
	// lang := "se"

	// file := "wordlist-german.txt"
	// lang := "de"

	file := "dictionary-pl.txt"
	lang := "pl"

	records, err := ReadTxtRecords(file)
	if err != nil {
		fmt.Println("Error reading records: " + err.Error())
		return
	}
	n := 2000
	top_n_records := GetTopNRecords(records, n)
	wordLengths := []int{5, 6, 7, 8}
	for _, length := range wordLengths {
		SaveAnswers(top_n_records, length, lang)
	}
}

func ReadTxtRecords(fileName string) ([][]string, error) {
	fmt.Println("Parsing answers from file '" + fileName + "'...")

	data, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error reading file: " + err.Error())
		return nil, err
	}

	scanner := bufio.NewScanner(data)
	var records [][]string
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

func SaveAnswers(records [][]string, length int, lang string) {
	langCapitalized := strings.ToUpper(lang[0:1]) + lang[1:]
	fileName := fmt.Sprintf("answers_%s_%d.go", lang, length)
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file: " + err.Error())
		return
	}

	defer f.Close()

	words := GetWordsOfLength(records, length, lang)

	fmt.Fprintln(f, "package answers")
	fmt.Fprintf(f, "var Answers%s%d = []string{\n", langCapitalized, length)
	for _, word := range words {
		text := fmt.Sprintf("\"%s\",\n", word)
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
