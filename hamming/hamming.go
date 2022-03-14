package hamming

import "errors"

func Distance(a, b string) (int, error) {

	count := 0
	if len(a) == len(b) {

		for index, runeValue := range a {
			if string(runeValue) != string(b[index]) {
				count++
			}
		}
		return count, nil
	}
	return 0, errors.New("error occured")
}
