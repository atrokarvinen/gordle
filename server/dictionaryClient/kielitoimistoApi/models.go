package kielitoimistoApi

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
}

type Response struct {
	Senses []Sense `json:"senses"`
	Word   string  `json:"clean_headword"`
}

type Sense struct {
	Examples    []Example    `json:"examples"`
	Definitions []Definition `json:"definitions"`
	Senses      []Sense      `json:"senses"`
}

type Example struct {
	Text string `json:"text"`
}

type Definition struct {
	Definition string `json:"definition"`
}
