package api

import (
	"go-test/game"
	"go-test/user"
	"go-test/wordsApi"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Api struct {
	DataProvider game.DataProvider
	Game         game.Game
	User         user.User
	WordsApi     wordsApi.IWordsApiClient
}

func (a Api) Run() {
	host_address := os.Getenv("HOST_ADDRESS")
	client_address := os.Getenv("CLIENT_ADDRESS")

	mode := os.Getenv("GIN_MODE")
	gin.SetMode(mode)
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{client_address},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/games", a.GetGames)
	r.GET("/games/latest", a.GetLatestGame)
	r.GET("/games/:id", a.GetGame)
	r.POST("/games", a.CreateGame)
	r.PUT("/games/:id", a.UpdateGame)
	r.DELETE("/games/:id", a.DeleteGame)

	r.POST("/games/:id/guesses", a.GuessWord)

	r.GET("users/new", a.CreateNewUser)
	r.GET("users/me", a.GetUserExists)
	r.POST("users/login", a.Login)
	r.DELETE("users/logout", a.Logout)

	r.GET("/echo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	r.Run(host_address)
}
