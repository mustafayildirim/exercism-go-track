package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {

	wc := make(Frequency)
	words := strings.FieldsFunc(phrase, func(r rune) bool {
		return !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\'')
	})
	for _, w := range words {
		w = strings.Trim(strings.ToLower(w), "'")
		wc[w] += 1
	}
	return wc
}
