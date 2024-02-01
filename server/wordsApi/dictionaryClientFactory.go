package wordsApi

type DictionaryClientFactory struct {
	DictionaryClientEn WordsApiClient
	DictionaryClientFi DictClientFi
}

func (d DictionaryClientFactory) GetDictionaryClient(language string) IWordsApiClient {
	if language == "en" {
		return d.DictionaryClientEn
	} else if language == "fi" {
		return d.DictionaryClientFi
	}
	panic("Invalid language '" + language + "'.")
}
