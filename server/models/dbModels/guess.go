package dbModels

import "gorm.io/gorm"

type Guess struct {
	gorm.Model
	GameID int
	Word   string
}
