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
		if difficulty == "all_words" {
			return AnswersFi5
		}
		return AnswersFiEasy5
	case 6:
		if difficulty == "all_words" {
			return AnswersFi6
		}
		return AnswersFiEasy6
	case 7:
		if difficulty == "all_words" {
			return AnswersFi7
		}
		return AnswersFiEasy7
	case 8:
		if difficulty == "all_words" {
			return AnswersFi8
		}
		return AnswersFiEasy8
	}
	panic(fmt.Sprintf("Unknown word length '%d'", wordLength))
}

func GetAnswersSe(difficulty string, wordLength int) []string {
	switch wordLength {
	case 5:
		if difficulty == "all_words" {
			return AnswersSe5
		}
		return AnswersSeEasy5
	case 6:
		if difficulty == "all_words" {
			return AnswersSe6
		}
		return AnswersSeEasy6
	case 7:
		if difficulty == "all_words" {
			return AnswersSe7
		}
		return AnswersSeEasy7
	case 8:
		if difficulty == "all_words" {
			return AnswersSe8
		}
		return AnswersSeEasy8
	}
	panic(fmt.Sprintf("Unknown word length '%d'", wordLength))
}

func GetAnswersDe(difficulty string, wordLength int) []string {
	switch wordLength {
	case 5:
		if difficulty == "all_words" {
			return AnswersDe5
		}
		return AnswersDeEasy5
	case 6:
		if difficulty == "all_words" {
			return AnswersDe6
		}
		return AnswersDeEasy6
	case 7:
		if difficulty == "all_words" {
			return AnswersDe7
		}
		return AnswersDeEasy7
	case 8:
		if difficulty == "all_words" {
			return AnswersDe8
		}
		return AnswersDeEasy8
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
