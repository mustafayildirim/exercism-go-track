package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {

	isbn = strings.ReplaceAll(isbn, "-", "")
	total := 0
	if len(isbn) != 10 {
		return false
	}

	for i := 0; i < 9; i++ {
		val, err := strconv.Atoi(string(isbn[i]))
		if err != nil {
			return false
		}
		total += (10 - i) * val
	}

	lastChar := string(isbn[9])
	if lastChar == "X" {
		total += 10
	} else {
		lastValue, err := strconv.Atoi(lastChar)
		if err != nil {
			return false
		} else {
			total += lastValue
		}
	}

	res := total % 11
	if res == 0 {
		return true
	}
	return false
}
