package luhn

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

func Valid(id string) bool {

	card := strings.Replace(id, " ", "", -1)

	cardLength := len(card)
	if cardLength < 2 {
		return false
	}

	doubleIndex := 1
	if cardLength%2 == 0 {
		doubleIndex = 0
	}

	total := 0
	for i, w := 0, 0; i < cardLength; i += w {
		runeValue, width := utf8.DecodeRuneInString(card[i:])
		w = width
		number, err := strconv.Atoi(string(runeValue))
		if err != nil {
			return false
		}
		if (i+doubleIndex)%2 == 0 {
			number = number * 2
			if number > 9 {
				number -= 9
			}
		}

		total += number
	}

	return total%10 == 0
}
