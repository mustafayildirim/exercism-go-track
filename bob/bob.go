package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {

	remark = strings.Replace(remark, "\t", "", -1)
	remark = strings.Replace(remark, "\r", "", -1)
	remark = strings.Replace(remark, "\n", "", -1)
	remark = strings.Trim(remark, " ")

	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	lastChar := remark[len(remark)-1:]
	if lastChar == "?" {
		if IsUpper(remark) && HasLetter(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	} else if lastChar == "!" {
		if IsUpper(remark) {
			return "Whoa, chill out!"
		}
	} else if !HasLetter(remark) {
		return "Whatever."
	} else {

		if IsUpper(remark) && HasLetter(remark) {
			return "Whoa, chill out!"
		}
	}

	return "Whatever."
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func HasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}

	return false
}
