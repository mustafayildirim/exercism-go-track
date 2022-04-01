package etl

import "strings"

func Transform(in map[int][]string) map[string]int {

	transformed := map[string]int{}

	for key, value := range in {
		for j := 0; j < len(value); j++ {

			transformed[strings.ToLower(value[j])] = key
		}

	}

	return transformed
}
