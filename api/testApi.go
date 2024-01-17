package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "John"},
	{ID: 2, Name: "Doe"},
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	iid, _ := strconv.Atoi(id)
	for _, user := range users {
		if user.ID == iid {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	users = append(users, user)
	c.JSON(http.StatusOK, users)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	iid, _ := strconv.Atoi(id)
	var user User
	c.BindJSON(&user)
	users[iid] = user
	c.JSON(http.StatusOK, users)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	iid, _ := strconv.Atoi(id)
	users = append(users[:iid], users[iid+1:]...)
	c.JSON(http.StatusOK, users)
}

func Test() {
	code := http.StatusOK
	fmt.Println(code)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/users", GetUsers)
	r.GET("/users/:id", GetUser)
	r.POST("/users", CreateUser)
	r.PUT("/users/:id", UpdateUser)
	r.DELETE("/users/:id", DeleteUser)

	r.Run("localhost:9000") // listen and serve on 0.0.0.0:8080
}
