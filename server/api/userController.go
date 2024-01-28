package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const userIdCookie = "user_id"

type LoginRequest struct {
	UserId int `json:"userId"`
}

func (a Api) GetUser(c *gin.Context) {
	userIdStr, err := c.Cookie(userIdCookie)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	user, err := a.User.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (a Api) Login(c *gin.Context) {
	var loginRequest LoginRequest
	err := c.BindJSON(&loginRequest)
	fmt.Println("LoginRequest.UserId:", loginRequest.UserId, "loginRequest.UserId == 0 => ", loginRequest.UserId == 0)
	if err != nil || loginRequest.UserId == 0 {
		fmt.Println("Failed to bind body to LoginRequest, creating new user...")

		user := a.User.CreateUser()

		fmt.Println("Created user:", user)

		userId := fmt.Sprint(user.ID)
		c.SetCookie(userIdCookie, userId, 3600, "/", "", false, true)
		c.JSON(http.StatusOK, user)
		return
	}
	userId := loginRequest.UserId
	user, err := a.User.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	userIdStr := fmt.Sprint(user.ID)
	c.SetCookie(userIdCookie, userIdStr, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, user)
}

func (a Api) Logout(c *gin.Context) {
	c.SetCookie(userIdCookie, "", -1, "/", "", false, true)
	c.Status(http.StatusOK)
}
