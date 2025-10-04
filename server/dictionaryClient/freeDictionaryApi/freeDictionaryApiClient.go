package freeDictionaryApi

import (
	"encoding/json"
	"gordle/dictionaryClient/models"
	"net/http"
	"net/url"
)

type FreeDictionaryApiClient struct{}

func (d FreeDictionaryApiClient) GetWord(word string, lang string) (models.DictionaryDetails, error) {
	urlEncodedWord := url.QueryEscape(word)
	url := "https://freedictionaryapi.com/api/v1/entries/" + lang + "/" + urlEncodedWord

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.DictionaryDetails{}, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return models.DictionaryDetails{}, err
	}
	defer response.Body.Close()

	var details Response
	err = json.NewDecoder(response.Body).Decode(&details)
	if err != nil {
		return models.DictionaryDetails{}, err
	}
	mappedDetails := mapWordDetails(details)

	return mappedDetails, nil
}
