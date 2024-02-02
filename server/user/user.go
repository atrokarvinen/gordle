package user

import (
	"fmt"
	"gordle/database"
	"gordle/models/dbModels"
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
	fmt.Println("Creating new user...")
	createdUser := u.Db.CreateUser()
	fmt.Println("Created user with id:", createdUser.ID)
	return createdUser
}

func (u User) DeleteUser(userId int) {
	u.Db.DeleteUser(userId)
}
