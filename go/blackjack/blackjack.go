package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	v := 0
	switch card {
	case "one":
		v = 1
	case "two":
		v = 2
	case "three":
		v = 3
	case "four":
		v = 4
	case "five":
		v = 5
	case "six":
		v = 6
	case "seven":
		v = 7
	case "eight":
		v = 8
	case "nine":
		v = 9
	case "ten", "jack", "queen", "king":
		v = 10
	case "ace":
		v = 11
	}
	return v
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	hand, dc := ParseCard(card1)+ParseCard(card2), ParseCard(dealerCard)
	switch {
	case hand == 22:
		return "P"
	case hand == 21:
		if dc < 10 {
			return "W"
		}
		return "S"
	case hand >= 17:
		return "S"
	case hand >= 12:
		if dc >= 7 {
			return "H"
		}
		return "S"
	default:
		return "H"
	}
}
