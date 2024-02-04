package game

import (
	"reflect"
	"testing"
)

func TestGetLeastPlayedAnswers(t *testing.T) {
	var tests = []struct {
		name          string
		playedAnswers []string
		allAnswers    []string
		expect        []string
	}{
		{
			name:          "returns all answers when playedAnswers is empty",
			playedAnswers: []string{},
			allAnswers:    []string{"apple", "banana"},
			expect:        []string{"apple", "banana"},
		},
		{
			name:          "returns least played answer when others have been played two rounds",
			playedAnswers: []string{"apple", "apple", "banana", "banana"},
			allAnswers:    []string{"apple", "banana", "cherry"},
			expect:        []string{"cherry"},
		},
		{
			name:          "returns least played answer when all have been played once",
			playedAnswers: []string{"apple", "banana", "banana"},
			allAnswers:    []string{"apple", "banana"},
			expect:        []string{"apple"},
		},
		{
			name:          "returns correct answer with scandic letters",
			playedAnswers: []string{"löyly"},
			allAnswers:    []string{"löyly", "höylä"},
			expect:        []string{"höylä"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := GetLeastPlayedAnswers(test.playedAnswers, test.allAnswers)
			if !reflect.DeepEqual(result, test.expect) {
				t.Errorf("GetLeastPlayedAnswers(%q, %q) = %q; want %q", test.playedAnswers, test.allAnswers, result, test.expect)
			}
		})
	}
}
