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
