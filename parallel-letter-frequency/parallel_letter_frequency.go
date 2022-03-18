package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	queue := make(chan FreqMap, len(l))
	for _, s := range l {
		queue <- Frequency(s)
	}
	close(queue)
	combined := FreqMap{}
	for elem := range queue {
		for k, v := range elem {
			combined[k] = v + combined[k]
		}
	}

	return combined
}
