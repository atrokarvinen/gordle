package dictionaryClient

import "go-test/dictionaryClient/models"

type IDictionaryClient interface {
	GetWord(string) (models.DictionaryDetails, error)
}
