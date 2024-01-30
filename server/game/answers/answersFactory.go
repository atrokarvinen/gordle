package answers

import "fmt"

func GetAnswers(lang string, wordLength int) []string {
	if lang == "fi" {
		return GetAnswersFi(wordLength)
	} else if lang == "en" {
		return GetAnswersEn(wordLength)
	}
	panic(fmt.Sprintf("Unknown language '%s'", lang))
}

func GetAnswersEn(wordLength int) []string {
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

func GetAnswersFi(wordLength int) []string {
	if wordLength == 5 {
		return AnswersFi5
	} else if wordLength == 6 {
		return AnswersFi6
	} else if wordLength == 7 {
		return AnswersFi7
	} else if wordLength == 8 {
		return AnswersFi8
	}
	panic(fmt.Sprintf("Unknown word length '%d'", wordLength))
}
