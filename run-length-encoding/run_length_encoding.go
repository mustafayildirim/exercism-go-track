package encode

import (
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {
	compressedString := ""
	if len(input) == 0 {
		return compressedString
	}
	charCount := 0
	for i, c := range input {
		charCount++
		// We won't overflow since
		// the runtime does not evaluate the
		// right-hand side unless the
		// left-hand side is falsy, in which case
		// the right-hand side would be safe to evaluate
		if i == len(input)-1 || c != rune(input[i+1]) {
			if charCount <= 1 {
				compressedString += string(c)
			} else {
				compressedString += strconv.Itoa(charCount) + string(c)
			}
			charCount = 0
		}
	}
	return compressedString
}
func RunLengthDecode(input string) string {
	decodedString := ""
	if len(input) == 0 {
		return decodedString
	}
	// If the string does not contain any numbers just return it as-is
	if !strings.ContainsAny(input, "1234567890") {
		return input
	}
	repeatNum := ""
	for _, c := range input {
		switch {
		// scan for numbers
		case unicode.IsDigit(c):
			repeatNum += string(c)
		// we found a letter and we have stored numbers,
		// so we must repeat the letter
		case len(repeatNum) > 0:
			r, _ := strconv.Atoi(repeatNum)
			decodedString += strings.Repeat(string(c), r)
			repeatNum = ""
		// we found a letter but we haven't got any number stored,
		// so it must be a single letter
		default:
			decodedString += string(c)
		}
	}
	return decodedString
}
