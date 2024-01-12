package game

import (
	"strings"
	"testing"
)

func TestCheckWord(t *testing.T) {
	result := CheckWord("gordle", "gordle")
	resultString := strings.Join(result, "")
	if resultString != "vvvvvv" {
		t.Errorf("CheckWord('gordle', 'gordle') = %q; want 'vvvvvv'", resultString)
	}
}

func TestCheckWord_FirstCorrect(t *testing.T) {
	result := CheckWord("gppppp", "gordle")
	resultString := strings.Join(result, "")
	if resultString != "vxxxxx" {
		t.Errorf("CheckWord('gppppp', 'gordle') = %q; want 'vxxxxx'", resultString)
	}
}

func TestCheckWord_ContainsButInWrongPlace(t *testing.T) {
	result := CheckWord("agaaaa", "gordle")
	resultString := strings.Join(result, "")
	if resultString != "x?xxxx" {
		t.Errorf("CheckWord('agaaaa', 'gordle') = %q; want 'x?xxxx'", resultString)
	}
}

func TestCheckWord_ContainsOnlyOne(t *testing.T) {
	result := CheckWord("aggggg", "gordle")
	resultString := strings.Join(result, "")
	if resultString != "x?xxxx" {
		t.Errorf("CheckWord('aggggg', 'gordle') = %q; want 'x?xxxx'", resultString)
	}
}

func TestCheckWord_ContainsNoneWhenCorrectShown(t *testing.T) {
	result := CheckWord("gggggg", "gordle")
	resultString := strings.Join(result, "")
	if resultString != "vxxxxx" {
		t.Errorf("CheckWord('gggggg', 'gordle') = %q; want 'vxxxxx'", resultString)
	}
}

func TestCheckWord_ContainsNoneWhenCorrectExists(t *testing.T) {
	result := CheckWord("eeeeee", "gordle")
	resultString := strings.Join(result, "")
	if resultString != "xxxxxv" {
		t.Errorf("CheckWord('eeeeee', 'gordle') = %q; want 'xxxxxv'", resultString)
	}
}

func TestCheckWord_AllWrong(t *testing.T) {
	result := CheckWord("qwtyui", "gordle")
	resultString := strings.Join(result, "")
	if resultString != "xxxxxx" {
		t.Errorf("CheckWord('qwtyui', 'gordle') = %q; want 'xxxxxx'", resultString)
	}
}

func TestCheckWord_AllContained(t *testing.T) {
	result := CheckWord("eldrog", "gordle")
	resultString := strings.Join(result, "")
	if resultString != "??????" {
		t.Errorf("CheckWord('eldrog', 'gordle') = %q; want '??????'", resultString)
	}
}
