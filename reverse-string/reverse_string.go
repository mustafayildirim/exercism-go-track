package reverse

import "unicode/utf8"

func Reverse(input string) string {
	output := ""
	for i, w := 0, 0; i < len(input); i += w {
		runeValue, width := utf8.DecodeRuneInString(input[i:])
		w = width
		output = string(runeValue) + output

	}

	return output
}
