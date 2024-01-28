package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const cookieName = "user_id"

type LoginRequest struct {
	UserId int `json:"userId"`
}

func (a Api) CreateNewUser(c *gin.Context) {
	user := a.User.CreateUser()
	fmt.Println("Created user:", user)
	userId := fmt.Sprint(user.ID)
	c.SetCookie(cookieName, userId, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, user)
}

func (a Api) GetUserExists(c *gin.Context) {
	userIdStr, err := c.Cookie(cookieName)
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
		fmt.Println("Failed to bind body to LoginRequest")
		a.CreateNewUser(c)
		return
	}
	userId := loginRequest.UserId
	user, err := a.User.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	userIdStr := fmt.Sprint(user.ID)
	c.SetCookie(cookieName, userIdStr, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, user)
}

func (a Api) Logout(c *gin.Context) {
	c.SetCookie(cookieName, "", -1, "/", "", false, true)
	c.Status(http.StatusOK)
}
