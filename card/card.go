package card

import (
	"github.com/l3njo/play/misc"
)

// GetRandomCard returns a random card as a string.
func GetRandomCard() string {
	randomCard := &misc.Card{Face: misc.Faces[misc.Int(len(misc.Faces))], Suit: misc.Suits[misc.Int(len(misc.Suits))]}
	return misc.GetCard(randomCard)
}

// GetRandomCards returns a slice containing n non-unique cards as strings.
// For a unique set, use GetDeck()
// Returns an error if n <= 0.
func GetRandomCards(n int) (randomCards []string, err string) {
	randomCards = []string{}
	err = ""
	if n <= 0 {
		err = "Invalid number of cards"
		return
	}
	for len(randomCards) < n {
		randomCards = append(randomCards, GetRandomCard())
	}
	return
}
