package card

import (
	"math/rand"
)

func init() {
	// rand.Seed(time.Now().UTC().UnixNano())
}

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
	deck := GetOrderedDeck()
	for len(deck) > 0 {
		i := rand.Intn(len(deck))
		card := deck[i]
		deck[i] = deck[len(deck)-1]
		deck = deck[:len(deck)-1]
		shuffledDeck = append(shuffledDeck, card)
	}
	return
}

// GetRandomDeck returns a random deck of predefined size
func GetRandomDeck(size int) (randomDeck []string) {
	randomDeck = []string{}
	deck := GetShuffledDeck()
	for len(randomDeck) < size {
		i := rand.Intn(len(deck))
		card := deck[i]
		deck[i] = deck[len(deck)-1]
		deck = deck[:len(deck)-1]
		randomDeck = append(randomDeck, card)
	}
	return
}

// Shuffle
