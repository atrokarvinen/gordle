package game

type DataProvider interface {
	GetPreviousGuesses() []string
	AddGuess(string)
	ResetGuesses()
}

type InMemoryDataProvider struct {
}

var inMemoryGuesses []string

func (d InMemoryDataProvider) GetPreviousGuesses() []string {
	return inMemoryGuesses
}

func (d InMemoryDataProvider) AddGuess(guess string) {
	inMemoryGuesses = append(inMemoryGuesses, guess)
}

func (d InMemoryDataProvider) ResetGuesses() {
	inMemoryGuesses = []string{}
}

type DatabaseDataProvider struct {
}

func (d DatabaseDataProvider) GetPreviousGuesses() []string {
	return []string{"fromDb"}
}
