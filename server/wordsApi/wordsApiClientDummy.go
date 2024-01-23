package wordsApi

type WordsApiClientDummy struct{}

func (w WordsApiClientDummy) GetWord(word string) WordDetails {
	return WordDetails{}
}

func (w WordsApiClientDummy) WordExists(word string) bool {
	return true
}
