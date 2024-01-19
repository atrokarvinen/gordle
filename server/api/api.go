package api

import (
	"go-test/game"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Api struct {
	DataProvider game.DataProvider
	Game         game.Game
}

func (a Api) Run() {
	mode := os.Getenv("GIN_MODE")
	gin.SetMode(mode)
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/games", a.GetGames)
	r.GET("/games/latest", a.GetLatestGame)
	r.GET("/games/:id", a.GetGame)
	r.GET("/games/:id/guesses", a.GetGuesses)
	r.POST("/games", a.CreateGame)
	r.PUT("/games/:id", a.UpdateGame)
	r.DELETE("/games/:id", a.DeleteGame)

	r.POST("/games/:id/guesses", a.GuessWord)

	r.Run("localhost:9000")
}
