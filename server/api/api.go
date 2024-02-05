package api

import (
	"gordle/dictionaryClient"
	"gordle/game"
	"gordle/user"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Api struct {
	DataProvider      game.DataProvider
	Game              game.Game
	User              user.User
	DictionaryFactory dictionaryClient.DictionaryClientFactory
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
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/games/latest", a.GetLatestGame)
	r.GET("/games/:id", a.GetGame)
	r.POST("/games", a.CreateGame)
	r.DELETE("/games/:id", a.DeleteGame)

	r.POST("/games/:id/guesses", a.GuessWord)

	r.GET("users/me", a.GetUser)
	r.POST("users/login", a.Login)
	r.DELETE("users/logout", a.Logout)

	r.GET("/statistics", a.GetStatistics)

	r.PUT("/test/answer", a.SetAnswer)

	r.Run(host_address)
}
