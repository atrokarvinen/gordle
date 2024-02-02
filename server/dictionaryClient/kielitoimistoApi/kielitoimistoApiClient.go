package kielitoimistoApi

import (
	"encoding/json"
	"errors"
	"fmt"
	"gordle/dictionaryClient/models"
	"net/http"
	"net/url"
)

type KielitoimistoApiClient struct {
}

var AccessToken string
var RefreshToken string

func (w KielitoimistoApiClient) GetWord(word string) (models.DictionaryDetails, error) {
	accessToken, err := w.getAccessToken()
	if err != nil {
		return models.DictionaryDetails{}, err
	}
	details, err := w.getDetails(word, accessToken)
	if err != nil && err.Error() != "Unauthorized" {
		return models.DictionaryDetails{}, err
	}

	// Attempt to refresh token
	accessToken, err = w.refreshAccessToken()
	if err != nil {
		AccessToken = ""
		RefreshToken = ""
		return models.DictionaryDetails{}, err
	}
	details, err = w.getDetails(word, accessToken)
	if err != nil {
		return models.DictionaryDetails{}, err
	}
	return details, nil
}

func (w KielitoimistoApiClient) getDetails(word string, accessToken string) (models.DictionaryDetails, error) {
	urlEncodedWord := url.QueryEscape(word)
	url := "https://www.kielitoimistonsanakirja.fi/api/search/api/v1/search?keyword=" + urlEncodedWord
	fmt.Println("word url:", url)
	wordRequest, err := http.NewRequest("GET", url, nil)
	wordRequest.Header.Set("Authorization", "Bearer "+accessToken)
	if err != nil {
		fmt.Println("Error creating word request:", err.Error())
		return models.DictionaryDetails{}, err
	}

	wordRes, err := http.DefaultClient.Do(wordRequest)
	if err != nil {
		fmt.Println("Error getting word:", err.Error())
		return models.DictionaryDetails{}, err
	}
	defer wordRes.Body.Close()
	fmt.Println("word response status:", wordRes.StatusCode)
	if wordRes.StatusCode == http.StatusUnauthorized {
		return models.DictionaryDetails{}, errors.New("Unauthorized")
	}

	var responses []Response
	err = json.NewDecoder(wordRes.Body).Decode(&responses)
	if err != nil {
		fmt.Println("Error decoding word response:", err.Error())
		return models.DictionaryDetails{}, err
	}
	fmt.Println("word details:", responses)

	if len(responses) == 0 {
		return models.DictionaryDetails{}, nil
	}

	details := mapWordDetails(responses[0])

	fmt.Println("definitions:", details.Definitions)
	fmt.Println("examples:", details.Examples)

	return details, nil
}

func (w KielitoimistoApiClient) getAccessToken() (string, error) {
	if AccessToken != "" {
		fmt.Println("access token already exists in memory")
		return AccessToken, nil
	}
	loginUrl := "https://www.kielitoimistonsanakirja.fi/api/auth/anon-login"
	loginRequest, err := http.NewRequest("GET", loginUrl, nil)
	if err != nil {
		fmt.Println("Error creating login request:", err.Error())
		return "", err
	}
	loginRes, err := http.DefaultClient.Do(loginRequest)
	if err != nil {
		fmt.Println("Error logging in:", err.Error())
		return "", err
	}
	defer loginRes.Body.Close()

	var response LoginResponse
	err = json.NewDecoder(loginRes.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding login response:", err.Error())
		return "", err
	}
	fmt.Println("access token status:", loginRes.StatusCode)

	AccessToken = response.AccessToken
	RefreshToken = response.RefreshToken

	return response.AccessToken, nil
}

func (w KielitoimistoApiClient) refreshAccessToken() (string, error) {
	loginUrl := "https://www.kielitoimistonsanakirja.fi/api/auth/renew"
	loginRequest, err := http.NewRequest("GET", loginUrl, nil)
	loginRequest.Header.Set("Authorization", "Bearer "+RefreshToken)
	if err != nil {
		fmt.Println("Error creating refresh request:", err.Error())
		return "", err
	}
	loginRes, err := http.DefaultClient.Do(loginRequest)
	if err != nil {
		fmt.Println("Error refreshing token:", err.Error())
		return "", err
	}
	defer loginRes.Body.Close()

	var response RefreshResponse
	err = json.NewDecoder(loginRes.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding refresh response:", err.Error())
		return "", err
	}
	fmt.Println("access token refresh response status:", loginRes.StatusCode)

	AccessToken = response.AccessToken

	return response.AccessToken, nil
}
