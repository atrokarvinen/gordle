package api

import (
	"fmt"
	"go-test/models/dbModels"
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
	fmt.Println("LoginRequest.UserId:", loginRequest.UserId)
	if err != nil || loginRequest.UserId == 0 {
		fmt.Println("Failed to bind body to LoginRequest, creating new user...")

		user := a.User.CreateUser()
		SetLoginResponse(c, user)
		return
	}
	userId := loginRequest.UserId
	user, err := a.User.GetUser(userId)
	if err != nil {
		fmt.Println("Error finding user '", userId, "':", err.Error())
		user := a.User.CreateUser()
		SetLoginResponse(c, user)
		return
	}
	fmt.Println("Found User:", user)
	SetLoginResponse(c, user)
}

func SetLoginResponse(c *gin.Context, user dbModels.User) {
	userId := fmt.Sprint(user.ID)
	userIdStr := fmt.Sprint(userId)
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie(userIdCookie, userIdStr, 3600*24*7, "/", "", true, true)
	c.JSON(http.StatusOK, user)
}

func (a Api) Logout(c *gin.Context) {
	c.SetCookie(userIdCookie, "", -1, "/", "", false, true)
	c.Status(http.StatusOK)
}
