package wordsApi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type DictClientFi struct{}

func (w DictClientFi) GetWord(word string) (DictionaryDetails, error) {
	loginUrl := "https://www.kielitoimistonsanakirja.fi/api/auth/anon-login"
	loginRequest, err := http.NewRequest("GET", loginUrl, nil)
	if err != nil {
		fmt.Println("Error creating login request:", err.Error())
		return DictionaryDetails{}, err
	}
	loginRes, err := http.DefaultClient.Do(loginRequest)
	if err != nil {
		fmt.Println("Error logging in:", err.Error())
		return DictionaryDetails{}, err
	}
	defer loginRes.Body.Close()

	var message LoginResponse
	err = json.NewDecoder(loginRes.Body).Decode(&message)
	if err != nil {
		fmt.Println("Error decoding login response:", err.Error())
		return DictionaryDetails{}, err
	}
	fmt.Println("access token status:", loginRes.StatusCode)

	urlEncodedWord := url.QueryEscape(word)
	url := "https://www.kielitoimistonsanakirja.fi/api/search/api/v1/search?keyword=" + urlEncodedWord
	fmt.Println("word url:", url)
	wordRequest, err := http.NewRequest("GET", url, nil)
	wordRequest.Header.Set("Authorization", "Bearer "+message.AccessToken)
	if err != nil {
		fmt.Println("Error creating word request:", err.Error())
		return DictionaryDetails{}, err
	}

	wordRes, err := http.DefaultClient.Do(wordRequest)
	if err != nil {
		fmt.Println("Error getting word:", err.Error())
		return DictionaryDetails{}, err
	}
	defer wordRes.Body.Close()

	var responses []ResponseFi
	err = json.NewDecoder(wordRes.Body).Decode(&responses)
	if err != nil {
		fmt.Println("Error decoding word response:", err.Error())
		return DictionaryDetails{}, err
	}
	fmt.Println("word details:", responses)

	if len(responses) == 0 {
		return DictionaryDetails{}, nil
	}

	details := mapWordDetailsFi(responses[0])
	return details, nil
}

func mapWordDetailsFi(details ResponseFi) DictionaryDetails {
	return DictionaryDetails{
		Word:        details.Word,
		Definitions: mapDefinitionsFi(details.Senses),
		Examples:    mapExamplesFi(details.Senses),
	}
}

func mapDefinitionsFi(sense []SenseFi) []string {
	var result []string
	for _, sense := range sense {
		for _, definition := range sense.Definitions {
			sanitized := sanitizeDefinition(definition.Definition)
			fmt.Println("definition", definition, "=> sanitized:", sanitized)
			result = append(result, sanitized)
		}
	}
	return result
}

// Words may contain html like <span class="kt-pos">adv. </span>, filter out the html
func sanitizeDefinition(definition string) string {
	var sanitized []rune
	openingTagBegin := false
	openingTagEnd := false
	closingTagStart := false
	closingTagEnd := false
	for _, r := range definition {
		if r == '<' {
			closingTagStart = openingTagBegin
			openingTagBegin = true
		} else if r == '>' {
			closingTagEnd = openingTagEnd
			openingTagEnd = true
		} else if openingTagEnd && !closingTagStart {
			sanitized = append(sanitized, r)
		} else if closingTagStart && closingTagEnd {
			sanitized = append(sanitized, r)
		}
	}
	return string(sanitized)
}

func mapExamplesFi(senses []SenseFi) []string {
	var result []string
	for _, sense := range senses {
		for _, example := range sense.Examples {
			result = append(result, example.Text)
		}
	}
	return result
}

func (w DictClientFi) WordExists(word string) bool {
	_, err := w.GetWord(word)
	return err == nil
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type ResponseFi struct {
	Senses []SenseFi `json:"senses"`
	Word   string    `json:"clean_headword"`
}

type SenseFi struct {
	Examples    []ExampleFi    `json:"examples"`
	Definitions []DefinitionFi `json:"definitions"`
}

type ExampleFi struct {
	Text string `json:"text"`
}

type DefinitionFi struct {
	Definition string `json:"definition"`
}
