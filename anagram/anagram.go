package anagram

import (
	"reflect"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	lowerSubject := strings.ToLower(subject)
	anagramMap := getAnagramMap(lowerSubject)
	detected := []string{}

	for _, candidate := range candidates {
		lowerCandidate := strings.ToLower(candidate)
		if lowerCandidate == lowerSubject {
			continue
		}

		candidateMap := getAnagramMap(lowerCandidate)
		eq := reflect.DeepEqual(anagramMap, candidateMap)
		if eq {
			detected = append(detected, candidate)
		}
	}

	return detected
}

func getAnagramMap(subject string) map[rune]int {
	anagramMap := map[rune]int{}
	for _, v := range subject {
		value, present := anagramMap[v]
		if present {
			anagramMap[v] = value + 1
		} else {
			anagramMap[v] = 1
		}
	}

	return anagramMap
}
