package dbModels

import "gorm.io/gorm"

type WordsApiCall struct {
	gorm.Model
	Word      string
	CreatedAt string
}
