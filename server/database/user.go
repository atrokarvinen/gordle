package database

import (
	"fmt"
	"go-test/models/dbModels"
)

func (d DatabaseDataProvider) CreateUser() dbModels.User {
	user := dbModels.User{}
	d.Db.Create(&user)
	return user
}

func (d DatabaseDataProvider) GetUser(userId int) (dbModels.User, error) {
	fmt.Println("Getting user", userId)
	var user dbModels.User
	result := d.Db.First(&user, userId)
	return user, result.Error
}

func (d DatabaseDataProvider) DeleteUser(userId int) {
	d.Db.Delete(&dbModels.User{}, userId)
}
