package api

import (
	"fmt"
	"gordle/models"
	"gordle/models/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (a Api) GetGameHistory(c *gin.Context) {
	userId, err := getUserIdFromContext(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		return
	}
	games, totalCount := a.Game.GetGames(userId, page, limit)
	dto := dto.GameHistoryResponse{
		TotalCount: totalCount,
		Games:      games,
	}
	c.JSON(http.StatusOK, dto)
}

func (a Api) GetStatistics(c *gin.Context) {
	userId, err := getUserIdFromContext(c)
	if err != nil {
		return
	}
	games := a.Game.GetStatistics(userId)
	totalStats := getStatisticsFromGames(games)

	gamesByLanguage := groupGamesByLanguage(games)
	statsByLanguage := make(map[string]map[int]dto.WordLengthMap)
	for lang, games := range gamesByLanguage {
		gamesByWordLength := groupGamesByWordLength(games)
		for wordLength, games := range gamesByWordLength {
			wordLenMap := dto.WordLengthMap{}
			wordLenMap.Total = getStatisticsFromGames(games)
			wordLenMap.ByGuessCount = make(map[int]int)
			gamesBvGuessCount := groupGamesByGuessCount(games)
			for guessCount, games := range gamesBvGuessCount {
				wins := filterWinGames(games)
				wordLenMap.ByGuessCount[guessCount] = len(wins)
			}
			if _, ok := statsByLanguage[lang]; !ok {
				statsByLanguage[lang] = make(map[int]dto.WordLengthMap)
			}
			statsByLanguage[lang][wordLength] = wordLenMap
		}
	}

	dto := dto.StatisticsResponse{
		AllGames:   games,
		Total:      totalStats,
		ByLanguage: statsByLanguage,
	}
	c.JSON(http.StatusOK, dto)
}

func groupGamesByLanguage(games []models.Game) map[string][]models.Game {
	gamesByLanguage := make(map[string][]models.Game)
	for _, game := range games {
		if _, ok := gamesByLanguage[game.Language]; !ok {
			gamesByLanguage[game.Language] = []models.Game{}
		}
		gamesByLanguage[game.Language] = append(gamesByLanguage[game.Language], game)
	}
	return gamesByLanguage
}

func groupGamesByWordLength(games []models.Game) map[int][]models.Game {
	gamesByWordLength := make(map[int][]models.Game)
	for _, game := range games {
		if _, ok := gamesByWordLength[game.WordLength]; !ok {
			gamesByWordLength[game.WordLength] = []models.Game{}
		}
		gamesByWordLength[game.WordLength] = append(gamesByWordLength[game.WordLength], game)
	}
	return gamesByWordLength
}

func groupGamesByGuessCount(games []models.Game) map[int][]models.Game {
	winGames := filterWinGames(games)
	winsByGuessCount := make(map[int][]models.Game)
	for _, game := range winGames {
		fmt.Println("Game", game.Answer, "guess count: ", len(game.Guesses))
		if _, ok := winsByGuessCount[len(game.Guesses)]; !ok {
			winsByGuessCount[len(game.Guesses)] = []models.Game{}
		}
		winsByGuessCount[len(game.Guesses)] = append(winsByGuessCount[len(game.Guesses)], game)
	}
	return winsByGuessCount
}

func getStatisticsFromGames(games []models.Game) dto.Statistics {
	winGames := filterWinGames(games)
	winCount := len(winGames)
	stats := dto.Statistics{
		WinCount:   winCount,
		LossCount:  len(games) - winCount,
		TotalCount: len(games),
		WinRate:    float64(winCount) / float64(len(games)),
	}
	return stats
}

func filterWinGames(games []models.Game) []models.Game {
	winGames := []models.Game{}
	for _, game := range games {
		if game.State == 2 {
			winGames = append(winGames, game)
		}
	}
	return winGames
}
