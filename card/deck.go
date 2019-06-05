package card

// GetOrderedDeck returns an entire ordered deck as a slice
func GetOrderedDeck() (deck []string) {
	deck = []string{}
	for _, suit := range suits {
		for _, face := range faces {
			card := &Card{face: face, suit: suit}
			deck = append(deck, card.getCard())
		}
	}
	return
}

// GetShuffledDeck returns an entire shuffled deck as a slice
func GetShuffledDeck() (shuffledDeck []string) {
	shuffledDeck = []string{}
	deck := GetShuffledDeck()
	pool := make([]string, len(deck))
	copy(pool, deck)
	for len(shuffledDeck) < len(deck) {
		card := GetRandomCard()
		if !inArray(card, shuffledDeck) {
			shuffledDeck = append(shuffledDeck, card)
		}
	}
	return
}

func inArray(needle string, haystack []string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

// Shuffle
// Get Random Deck
