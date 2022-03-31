package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var filtered Ints
	for _, v := range i {
		if filter(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var filtered Ints
	for _, v := range i {
		if !filter(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var filtered Lists
	for _, v := range l {
		if filter(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var filtered Strings
	for _, v := range s {
		if filter(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
