package wordsApi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
)

type WordsApiClient struct{}

func (w WordsApiClient) GetWord(word string) (WordDetails, error) {
	apiKey := os.Getenv("RAPID_API_KEY")

	if apiKey == "" {
		return WordDetails{}, nil
	}

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
		fmt.Printf("Word not %q found\n", word)
		return WordDetails{}, errors.New("Word not found")
	}

	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		return WordDetails{}, err
	}
	fmt.Println("response:", string(dump))

	var details WordDetails
	json.NewDecoder(response.Body).Decode(&details)
	fmt.Println("word details:", details)
	fmt.Println("results count:", len(details.Results))
	if len(details.Results) > 0 {
		fmt.Println("word definition #1:", details.Results[0].Definition)
	}
	return details, nil
}

func (w WordsApiClient) WordExists(word string) bool {
	_, err := w.GetWord(word)
	return err == nil
}
