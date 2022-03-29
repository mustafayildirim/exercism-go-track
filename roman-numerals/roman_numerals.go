package romannumerals

import (
	"errors"
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNumeral

var allRomanNumerals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(input int) (string, error) {

	var result strings.Builder

	if input < 1 || input > 3000 {
		return "", errors.New("number should between 1 and 3000")
	}

	for _, numeral := range allRomanNumerals {
		for input >= numeral.Value {
			result.WriteString(numeral.Symbol)
			input -= numeral.Value
		}
	}

	return result.String(), nil
}
