package toolkit

import (
	"testing"
	"unicode/utf8"
)

func TestTools_RandomString(t *testing.T) {
	var tools Tools
	s := tools.RandomString(10)
	if utf8.RuneCountInString(s) != 10 {
		t.Error("wrong length")
	}
}
