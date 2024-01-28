package dbModels

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Games []Game
}
