package toolkit

import (
	"math/rand"
	"unicode/utf8"
)

const randomStringSource = "abcdefghigklmnopqrstuvwxyzABCDEFGHIGKLMNOPQRSTUVWXYZ0123456789_+"

// Tools is the type to instantiate this module
type Tools struct {
}

// RandomString return a string of random characters of length n use randomStringSource
// as source and accept UTF8 generate
func (t *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(randomStringSource)
	for i := range s {
		s[i] = r[rand.Intn(utf8.RuneCountInString(randomStringSource))]
	}
	return string(s)
}
