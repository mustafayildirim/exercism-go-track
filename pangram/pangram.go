package pangram

import "strings"

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	pangram := "abcdefghijklmnopqrstuvwxyz"
	for _, v := range pangram {
		if !strings.Contains(input, string(v)) {
			return false
		}
	}
	return true
}
