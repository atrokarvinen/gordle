package dictionaryApi

import (
	"encoding/json"
	"gordle/dictionaryClient/models"
	"net/http"
)

type DictionaryApiClient struct{}

func (d DictionaryApiClient) GetWord(word string) (models.DictionaryDetails, error) {
	url := "https://api.dictionaryapi.dev/api/v2/entries/en/" + word
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.DictionaryDetails{}, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return models.DictionaryDetails{}, err
	}
	defer response.Body.Close()

	var details []Response
	err = json.NewDecoder(response.Body).Decode(&details)
	if err != nil {
		return models.DictionaryDetails{}, err
	}
	mappedDetails := mapWordDetails(details)

	return mappedDetails, nil
}
