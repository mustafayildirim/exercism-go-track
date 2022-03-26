package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	lengthOfRhyme := len(rhyme)

	var proverb []string
	if lengthOfRhyme >= 1 {
		for i, r := range rhyme {
			// Last element
			if i+1 == lengthOfRhyme {
				proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
			} else {
				proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", r, rhyme[i+1]))
			}
		}

	}

	return proverb
}
