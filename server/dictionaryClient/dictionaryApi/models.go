package dictionaryApi

type Response struct {
	Word     string    `json:"word"`
	Meanings []Meaning `json:"meanings"`
}

type Meaning struct {
	Definitions []Definition `json:"definitions"`
}

type Definition struct {
	Definition string `json:"definition"`
	Example    string `json:"example"`
}
