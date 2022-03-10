package logs

import (
	"unicode/utf8"
)

const RECOMMENDATION = "recommendation"
const RecommendationUnicode = "\u2757"
const SEARCH = "search"
const SearchUnicode = "\U0001F50D"
const WEATHER = "weather"
const WeatherUnicode = "\u2600"

// Application identifies the application emitting the given log.
func Application(log string) string {

	recommendationRune, _ := utf8.DecodeRuneInString(RecommendationUnicode)
	searchRune, _ := utf8.DecodeRuneInString(SearchUnicode)
	weatherRune, _ := utf8.DecodeRuneInString(WeatherUnicode)

	for i, w := 0, 0; i < len(log); i += w {
		runeValue, width := utf8.DecodeRuneInString(log[i:])
		w = width

		if log == "executed search 🔍" {
			println(runeValue)
			println(searchRune)
		}

		if runeValue == recommendationRune {
			return RECOMMENDATION
		} else if runeValue == searchRune {
			return SEARCH
		} else if runeValue == weatherRune {
			return WEATHER
		}

	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {

	var r []rune
	for i, w := 0, 0; i < len(log); i += w {
		runeValue, width := utf8.DecodeRuneInString(log[i:])
		if runeValue == oldRune {
			runeValue = newRune
		}
		r = append(r, runeValue)
		w = width
	}

	return string(r)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {

	numberOfRunes := utf8.RuneCountInString(log)
	return numberOfRunes <= limit
}
