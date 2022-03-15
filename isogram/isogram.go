package isogram

import (
	"strings"
	"unicode/utf8"
)

const HyphenUnicode = "\u002D"
const SpaceUnicode = "\u0020"

func IsIsogram(word string) bool {
	if len(word) == 0 {
		return true
	}

	values := map[rune]int{}
	hyphenRune, _ := utf8.DecodeRuneInString(HyphenUnicode)
	spaceRune, _ := utf8.DecodeRuneInString(SpaceUnicode)
	wordUpper := strings.ToUpper(word)
	for i, w := 0, 0; i < len(wordUpper); i += w {
		runeValue, width := utf8.DecodeRuneInString(wordUpper[i:])
		w = width
		if runeValue == hyphenRune || runeValue == spaceRune {
			continue
		}

		if _, ok := values[runeValue]; ok {
			return false
		} else {
			values[runeValue] = 1
		}

	}

	return true
}
