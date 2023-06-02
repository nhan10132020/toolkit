package toolkit

import (
	"os"
	"path/filepath"
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

func TestTool_CreateDirIfNotExist(t *testing.T) {
	var testTool Tools
	absPath, _ := filepath.Abs("./myDir")

	err := testTool.CreateDirIfNotExist(absPath)
	if err != nil {
		t.Error(err)
	}

	err = testTool.CreateDirIfNotExist(absPath)
	if err != nil {
		t.Error(err)
	}

	_ = os.Remove(absPath)
}

var slugTests = []struct {
	name          string
	s             string
	expected      string
	errorExpected bool
}{
	{name: "valid string", s: "now is the time", expected: "now-is-the-time", errorExpected: false},
	{name: "empty string", s: "", expected: "", errorExpected: true},
	{name: "comple string", s: "Now Is Time! + fist & & ^123 ", expected: "now-is-time-fist-123", errorExpected: false},
	{name: "japanese string", s: "こんにちは世界", expected: "", errorExpected: true},
	{name: "japanese string and roman characters", s: "こんにちは世界 hello world &+123", expected: "hello-world-123", errorExpected: false},
}

func TestTool_Slugify(t *testing.T) {
	var testTool Tools

	for _, e := range slugTests {
		slug, err := testTool.Slugify(e.s)
		if err != nil && !e.errorExpected {
			t.Errorf("%s: error recevied when none expected %s", e.name, err.Error())
		}

		if !e.errorExpected && slug != e.expected {
			t.Errorf("%s: wrong slug returned; expected %s but got %s", e.name, e.expected, slug)
		}
	}
}
