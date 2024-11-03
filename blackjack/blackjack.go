package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "king", "queen", "jack", "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerSum := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)

	if playerSum == 22 {
		return "P"
	} else if playerSum == 21 && dealerValue != 11 && dealerValue != 10 {
		return "W"
	} else if (playerSum == 21 && dealerValue == 11) || (playerSum == 21 && dealerValue == 10) {
		return "S"
	} else if playerSum >= 17 && playerSum <= 20 {
		return "S"
	} else if playerSum >= 12 && playerSum <= 16 && dealerValue >= 7 {
		return "H"
	} else if playerSum >= 12 && playerSum <= 16 && dealerValue < 7 {
		return "S"
	}
	return "H"
}
