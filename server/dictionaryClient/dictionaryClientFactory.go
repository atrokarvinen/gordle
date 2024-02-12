package dictionaryClient

type DictionaryClientFactory struct {
	DictionaryClientEn IDictionaryClient
	DictionaryClientFi IDictionaryClient
}

func (d DictionaryClientFactory) GetDictionaryClient(language string) IDictionaryClient {
	if language == "en" {
		return d.DictionaryClientEn
	} else if language == "fi" {
		return d.DictionaryClientFi
	}
	panic("Invalid language '" + language + "'.")
}
