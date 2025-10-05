package answers

import "fmt"

func GetAnswers(lang string, difficulty string, wordLength int) []string {
	switch lang {
	case "fi":
		return GetAnswersFi(difficulty, wordLength)
	case "en":
		return GetAnswersEn(difficulty, wordLength)
	case "se":
		return GetAnswersSe(difficulty, wordLength)
	case "de":
		return GetAnswersDe(difficulty, wordLength)
	case "pl":
		return GetAnswersPl(difficulty, wordLength)
	}
	panic(fmt.Sprintf("Unknown language '%s'", lang))
}

func GetAnswersEn(difficulty string, wordLength int) []string {
	switch wordLength {
	case 5:
		return AnswersEn5
	case 6:
		return AnswersEn6
	case 7:
		return AnswersEn7
	case 8:
		return AnswersEn8
	}
	panic(fmt.Sprintf("Unknown word length '%d'", wordLength))
}

func GetAnswersFi(difficulty string, wordLength int) []string {
	switch wordLength {
	case 5:
		return AnswersFi5
	case 6:
		return AnswersFi6
	case 7:
		return AnswersFi7
	case 8:
		return AnswersFi8
	}
	panic(fmt.Sprintf("Unknown word length '%d'", wordLength))
}

func GetAnswersSe(difficulty string, wordLength int) []string {
	switch wordLength {
	case 5:
		return AnswersSe5
	case 6:
		return AnswersSe6
	case 7:
		return AnswersSe7
	case 8:
		return AnswersSe8
	}
	panic(fmt.Sprintf("Unknown word length '%d'", wordLength))
}

func GetAnswersDe(difficulty string, wordLength int) []string {
	switch wordLength {
	case 5:
		return AnswersDe5
	case 6:
		return AnswersDe6
	case 7:
		return AnswersDe7
	case 8:
		return AnswersDe8
	}
	panic(fmt.Sprintf("Unknown word length '%d'", wordLength))
}

func GetAnswersPl(difficulty string, wordLength int) []string {
	switch wordLength {
	case 5:
		return AnswersPl5
	case 6:
		return AnswersPl6
	case 7:
		return AnswersPl7
	case 8:
		return AnswersPl8
	}
	panic(fmt.Sprintf("Unknown word length '%d'", wordLength))
}
