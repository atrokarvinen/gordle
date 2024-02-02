package models

type DictionaryDetails struct {
	Word        string   `json:"word"`
	Definitions []string `json:"definitions"`
	Examples    []string `json:"examples"`
}

func GetDefaultWordDetails(word string) DictionaryDetails {
	return DictionaryDetails{
		Word:        word,
		Definitions: []string{},
		Examples:    []string{},
	}
}
