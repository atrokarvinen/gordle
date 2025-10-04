package freeDictionaryApi

type Response struct {
	Word    string  `json:"word"`
	Entries []Entry `json:"entries"`
}

type Entry struct {
	// Language       string   `json:"language"`
	// PartOfSpeech   string   `json:"partOfSpeech"`
	// Pronunciations []string `json:"pronunciations"`
	// Forms          []string `json:"forms"`
	Senses []Sense `json:"senses"`
	// Synonyms       []string `json:"synonyms"`
	// Antonyms       []string `json:"antonyms"`
}

type Sense struct {
	Definition string   `json:"definition"`
	Examples   []string `json:"examples"`
}
