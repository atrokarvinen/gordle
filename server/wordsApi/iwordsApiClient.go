package wordsApi

type IWordsApiClient interface {
	GetWord(string) (WordDetails, error)
	WordExists(string) bool
}
