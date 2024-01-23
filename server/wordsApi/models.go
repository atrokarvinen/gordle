package wordsApi

type WordDetails struct {
	Word          string            `json:"word"`
	Results       []WordResults     `json:"results"`
	Syllables     WordSyllables     `json:"syllables"`
	Pronunciation WordPronunciation `json:"pronunciation"`
	Frequency     float64           `json:"frequency"`
}

type WordResults struct {
	Definition   string   `json:"definition"`
	PartOfSpeech string   `json:"partOfSpeech"`
	Synonyms     []string `json:"synonyms"`
	Typeof       []string `json:"typeOf"`
	HasTypes     []string `json:"hasTypes"`
	Derivation   []string `json:"derivation"`
	Examples     []string `json:"examples"`
}

type WordPronunciation struct {
	All string `json:"all"`
}

type WordSyllables struct {
	Count int      `json:"count"`
	List  []string `json:"list"`
}
