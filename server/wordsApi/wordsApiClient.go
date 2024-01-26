package wordsApi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type WordsApiClient struct{}

func (w WordsApiClient) GetWord(word string) (WordDetails, error) {
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

	fmt.Println("response status:", response.StatusCode)

	if response.StatusCode == http.StatusNotFound {
		fmt.Printf("Word %q not found\n", word)
		return WordDetails{}, errors.New("Word not found")
	} else if response.StatusCode == http.StatusForbidden {
		return WordDetails{}, errors.New("Invalid Api Key")
	} else if response.StatusCode == http.StatusTooManyRequests {
		return WordDetails{}, errors.New("Too many requests")
	} else if response.StatusCode == http.StatusUnauthorized {
		return WordDetails{}, errors.New("No Api Key provided")
	} else if response.StatusCode != http.StatusOK {
		return WordDetails{}, errors.New("Unknown error")
	}

	var details WordDetails
	json.NewDecoder(response.Body).Decode(&details)
	fmt.Println("word details:", details)
	return details, nil
}

func (w WordsApiClient) WordExists(word string) bool {
	_, err := w.GetWord(word)
	return err == nil
}

func GetDefaultWordDetails(word string) WordDetails {
	var defaultWord WordDetails = WordDetails{
		Word: word,
		Results: []WordResults{
			{Definition: "No definition found"},
		},
	}
	return defaultWord
}
