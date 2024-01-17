package game

import (
	"strings"
	"testing"
)

var game = Game{}

func TestCheckWord(t *testing.T) {
	result := game.CheckWord("gordle", "gordle", 6)
	resultString := strings.Join(result, "")
	if resultString != "vvvvvv" {
		t.Errorf("CheckWord('gordle', 'gordle') = %q; want 'vvvvvv'", resultString)
	}
}

func TestCheckWord_FirstCorrect(t *testing.T) {
	result := game.CheckWord("gppppp", "gordle", 6)
	resultString := strings.Join(result, "")
	if resultString != "vxxxxx" {
		t.Errorf("CheckWord('gppppp', 'gordle') = %q; want 'vxxxxx'", resultString)
	}
}

func TestCheckWord_ContainsButInWrongPlace(t *testing.T) {
	result := game.CheckWord("agaaaa", "gordle", 6)
	resultString := strings.Join(result, "")
	if resultString != "x?xxxx" {
		t.Errorf("CheckWord('agaaaa', 'gordle') = %q; want 'x?xxxx'", resultString)
	}
}

func TestCheckWord_ContainsOnlyOne(t *testing.T) {
	result := game.CheckWord("aggggg", "gordle", 6)
	resultString := strings.Join(result, "")
	if resultString != "x?xxxx" {
		t.Errorf("CheckWord('aggggg', 'gordle') = %q; want 'x?xxxx'", resultString)
	}
}

func TestCheckWord_ContainsNoneWhenCorrectShown(t *testing.T) {
	result := game.CheckWord("gggggg", "gordle", 6)
	resultString := strings.Join(result, "")
	if resultString != "vxxxxx" {
		t.Errorf("CheckWord('gggggg', 'gordle') = %q; want 'vxxxxx'", resultString)
	}
}

func TestCheckWord_ContainsNoneWhenCorrectExists(t *testing.T) {
	result := game.CheckWord("eeeeee", "gordle", 6)
	resultString := strings.Join(result, "")
	if resultString != "xxxxxv" {
		t.Errorf("CheckWord('eeeeee', 'gordle') = %q; want 'xxxxxv'", resultString)
	}
}

func TestCheckWord_AllWrong(t *testing.T) {
	result := game.CheckWord("qwtyui", "gordle", 6)
	resultString := strings.Join(result, "")
	if resultString != "xxxxxx" {
		t.Errorf("CheckWord('qwtyui', 'gordle') = %q; want 'xxxxxx'", resultString)
	}
}

func TestCheckWord_AllContained(t *testing.T) {
	result := game.CheckWord("eldrog", "gordle", 6)
	resultString := strings.Join(result, "")
	if resultString != "??????" {
		t.Errorf("CheckWord('eldrog', 'gordle') = %q; want '??????'", resultString)
	}
}
