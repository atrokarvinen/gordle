package dictionaryClient

type DictionaryClientFactory struct {
	DictionaryClientEn      IDictionaryClient
	DictionaryClientFi      IDictionaryClient
	DictionaryClientGeneral IDictionaryClient
}

func (d DictionaryClientFactory) GetDictionaryClient(language string) IDictionaryClient {
	switch language {
	case "en":
		return d.DictionaryClientEn
	case "fi":
		return d.DictionaryClientFi
	case "de", "se", "pl":
		return d.DictionaryClientGeneral
	}
	panic("Invalid language '" + language + "'.")
}
