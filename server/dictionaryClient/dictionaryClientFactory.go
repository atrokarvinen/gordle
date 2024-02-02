package dictionaryClient

import (
	"gordle/dictionaryClient/kielitoimistoApi"
	"gordle/dictionaryClient/wordsApi"
)

type DictionaryClientFactory struct {
	DictionaryClientEn wordsApi.WordsApiClient
	DictionaryClientFi kielitoimistoApi.KielitoimistoApiClient
}

func (d DictionaryClientFactory) GetDictionaryClient(language string) IDictionaryClient {
	if language == "en" {
		return d.DictionaryClientEn
	} else if language == "fi" {
		return d.DictionaryClientFi
	}
	panic("Invalid language '" + language + "'.")
}
