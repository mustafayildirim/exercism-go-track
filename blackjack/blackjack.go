package blackjack

func getCardPoint(card string) int {
	m := make(map[string]int)
	m["ace"] = 11
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4
	m["five"] = 5
	m["six"] = 6
	m["seven"] = 7
	m["eight"] = 8
	m["nine"] = 9
	m["ten"] = 10
	m["jack"] = 10
	m["queen"] = 10
	m["king"] = 10

	if val, ok := m[card]; ok {
		return val
	}

	return 0

}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return getCardPoint(card)
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	card1Point := getCardPoint(card1)
	card2Point := getCardPoint(card2)

	cardTotal := card1Point + card2Point
	if cardTotal == 21 {
		return true
	}
	return false
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {

	if isBlackjack {
		if dealerScore == 11 || dealerScore == 10 {
			return "S"
		} else if dealerScore < 21 {
			return "W"
		}
	} else {
		if dealerScore == 11 {
			return "P"
		}
	}

	return "S"
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	if handScore >= 17 {
		return "S"
	} else if handScore <= 11 {
		return "H"
	} else {
		if 16 >= handScore && handScore >= 12 && dealerScore < 7 {

			return "S"
		}
	}
	return "H"
}
