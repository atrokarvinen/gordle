package user

import (
	"go-test/database"
	"go-test/models/dbModels"
)

type User struct {
	Db database.DatabaseDataProvider
}

func (u User) GetUsers() []User {
	var users []User
	return users
}

func (u User) GetUser(userId int) (dbModels.User, error) {
	return u.Db.GetUser(userId)
}

func (u User) CreateUser() dbModels.User {
	createdUser := u.Db.CreateUser()
	return createdUser
}

func (u User) DeleteUser(userId int) {
	u.Db.DeleteUser(userId)
}
