package dictionaryClient

import "gordle/dictionaryClient/models"

type IDictionaryClient interface {
	GetWord(string, string) (models.DictionaryDetails, error)
}
