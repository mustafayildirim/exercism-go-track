package acronym

import "strings"

import "regexp"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	result := ""

	words := regexp.MustCompile("[ _-]").Split(s, -1)
	if words != nil {
		for _, word := range words {
			if len(word) > 0 {
				result += strings.ToUpper(string(word[0]))
			}
		}
	}

	return result
}
