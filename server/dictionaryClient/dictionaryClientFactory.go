package dictionaryClient

type DictionaryClientFactory struct {
	DictionaryClientEn      IDictionaryClient
	DictionaryClientFi      IDictionaryClient
	DictionaryClientGeneral IDictionaryClient
}

func (d DictionaryClientFactory) GetDictionaryClient(language string, difficulty string) IDictionaryClient {
	switch language {
	case "en":
		return d.DictionaryClientEn
	case "fi":
		if difficulty == "easy" {
			return d.DictionaryClientGeneral
		}
		return d.DictionaryClientFi
	case "de", "se", "pl":
		return d.DictionaryClientGeneral
	}
	panic("Invalid language '" + language + "'.")
}
