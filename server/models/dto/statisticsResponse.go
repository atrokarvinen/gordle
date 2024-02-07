package dto

type StatisticsResponse struct {
	Total      Statistics                       `json:"total"`
	ByLanguage map[string]map[int]WordLengthMap `json:"byLanguage"`
}

type WordLengthMap struct {
	Total        Statistics  `json:"total"`
	ByGuessCount map[int]int `json:"byGuessCount"`
}

type Statistics struct {
	WinCount   int     `json:"winCount"`
	LossCount  int     `json:"lossCount"`
	TotalCount int     `json:"totalCount"`
	WinRate    float64 `json:"winRate"`
}
