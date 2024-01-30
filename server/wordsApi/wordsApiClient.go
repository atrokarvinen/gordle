package wordsApi

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-test/database"
	"go-test/models/dbModels"
	"net/http"
	"os"
	"time"
)

type WordsApiClient struct {
	DataProvider database.DatabaseDataProvider
}

func (w WordsApiClient) GetWord(word string) (WordDetails, error) {
	fmt.Println("Getting word from Words API:", word)
	err := w.getApiCallsCount()
	if err != nil {
		return WordDetails{}, err
	}

	apiKey := os.Getenv("RAPID_API_KEY")
	url := "https://wordsapiv1.p.rapidapi.com/words/" + word
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-RapidAPI-Key", apiKey)
	req.Header.Set("X-RapidAPI-Host", "wordsapiv1.p.rapidapi.com")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return WordDetails{}, err
	}
	defer response.Body.Close()

	statusCode := response.StatusCode
	fmt.Println("response status:", statusCode)
	w.addApiCallToDb(word, statusCode)

	if statusCode == http.StatusNotFound {
		fmt.Printf("Word %q not found\n", word)
		return WordDetails{}, errors.New("Word '" + word + "' not found")
	} else if statusCode == http.StatusForbidden {
		return WordDetails{}, errors.New("Invalid Api Key")
	} else if statusCode == http.StatusTooManyRequests {
		return WordDetails{}, errors.New("Too many requests")
	} else if statusCode == http.StatusUnauthorized {
		return WordDetails{}, errors.New("No Api Key provided")
	} else if statusCode != http.StatusOK {
		return WordDetails{}, errors.New("Unknown error")
	}

	var details WordDetails
	decodeError := json.NewDecoder(response.Body).Decode(&details)
	if decodeError != nil {
		return WordDetails{}, decodeError
	}
	// fmt.Println("word details:", details)
	fmt.Printf("Word %q is valid\n", word)
	return details, nil
}

func (w WordsApiClient) WordExists(word string) bool {
	_, err := w.GetWord(word)
	return err == nil
}

func GetDefaultWordDetails(word string) WordDetails {
	var defaultWord WordDetails = WordDetails{
		Word:    word,
		Results: []WordResults{{Definition: ""}},
	}
	return defaultWord
}

func (w WordsApiClient) getApiCallsCount() error {
	// fmt.Println("Getting calls count...")

	calls := w.DataProvider.GetWordsApiCalls()

	minDate := time.Now().UTC().AddDate(0, 0, -1).String()
	pastCalls, dailyCalls := filterCallsByDate(calls, minDate)

	// fmt.Println("Total calls count:", len(calls))
	fmt.Println("Past 24h calls count:", len(dailyCalls))
	// fmt.Println("Beoynd 24h calls count:", len(pastCalls))

	if len(pastCalls) > 0 {
		w.DataProvider.DeleteWordsApiCalls(pastCalls)
	}

	dailyMax := 2500
	if len(dailyCalls) > dailyMax {
		return errors.New("Daily limit reached")
	}
	return nil
}

func filterCallsByDate(calls []dbModels.WordsApiCall, date string) ([]dbModels.WordsApiCall, []dbModels.WordsApiCall) {
	var before []dbModels.WordsApiCall
	var after []dbModels.WordsApiCall
	for _, call := range calls {
		if date > call.CreatedAt {
			before = append(before, call)
		}
		if call.CreatedAt > date {
			after = append(after, call)
		}
	}
	return before, after
}

func (w WordsApiClient) addApiCallToDb(word string, statusCode int) {
	apiCallRegistered := statusCode == http.StatusOK || statusCode == http.StatusNotFound
	if !apiCallRegistered {
		return
	}

	now := time.Now().UTC().String()
	call := dbModels.WordsApiCall{Word: word, CreatedAt: now}

	fmt.Println("Adding api call...")

	w.DataProvider.AddWordsApiCall(call)
}
