package dictionaryClient

type DictionaryClientFactory struct {
	DictionaryClientEn IDictionaryClient
	DictionaryClientFi IDictionaryClient
}

func (d DictionaryClientFactory) GetDictionaryClient(language string) IDictionaryClient {
	switch language {
	case "en":
		return d.DictionaryClientEn
	case "fi":
		return d.DictionaryClientFi
	}
	panic("Invalid language '" + language + "'.")
}
