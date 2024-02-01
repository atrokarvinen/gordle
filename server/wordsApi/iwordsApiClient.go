package wordsApi

type IWordsApiClient interface {
	GetWord(string) (DictionaryDetails, error)
	WordExists(string) bool
}
