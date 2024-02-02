package dictionaryClient

import "gordle/dictionaryClient/models"

type IDictionaryClient interface {
	GetWord(string) (models.DictionaryDetails, error)
}
