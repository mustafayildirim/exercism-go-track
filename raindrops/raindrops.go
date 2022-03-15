package raindrops

import "strconv"

func Convert(number int) string {
	convertedValue := ""

	if number%3 == 0 {
		convertedValue += "Pling"
	}

	if number%5 == 0 {
		convertedValue += "Plang"
	}

	if number%7 == 0 {
		convertedValue += "Plong"
	}

	if convertedValue == "" {
		convertedValue = strconv.Itoa(number)
	}

	return convertedValue
}
