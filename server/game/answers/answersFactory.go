package answers

import "fmt"

func GetAnswers(wordLength int) []string {
	if wordLength == 5 {
		return AnswersLetters5
	} else if wordLength == 6 {
		return AnswersLetters6
	} else if wordLength == 7 {
		return AnswersLetters7
	} else if wordLength == 8 {
		return AnswersLetters8
	}
	panic(fmt.Sprintf("Unknown word length '%d'", wordLength))
}
