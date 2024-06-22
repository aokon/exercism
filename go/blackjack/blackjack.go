package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var result int

	switch card {
	case "ace":
		result = 11
	case "two":
		result = 2
	case "three":
		result = 3
	case "four":
		result = 4
	case "five":
		result = 5
	case "six":
		result = 6
	case "seven":
		result = 7
	case "eight":
		result = 8
	case "nine":
		result = 9
	case "ten":
		result = 10
	case "jack":
		result = 10
	case "queen":
		result = 10
	case "king":
		result = 10
	}

	return result
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var card1Num, card2Num, dealerCardNum = ParseCard(card1), ParseCard(card2), ParseCard(dealerCard)
	var result string

	cardSum := card1Num + card2Num

	switch {
	case card1 == "ace" && card2 == "ace":
		result = "P"
	case cardSum <= 11:
		result = "H"
	case cardSum >= 17 && cardSum <= 20:
		result = "S"
	case (cardSum >= 12 && cardSum <= 16) && dealerCardNum >= 7:
		result = "H"
	case (cardSum >= 12 && cardSum <= 16):
		result = "S"
	case cardSum == 21 && (dealerCard == "ace" || dealerCard == "ten" || dealerCard == "queen" || dealerCard == "king"):
		result = "S"
	default:
		result = "W"
	}

	return result
}
