package card

import (
	"math/rand"
)

type card struct {
	face, suit string
}

// Card face value
var faces = [...]string{"Ace", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
var suits = [...]string{"Spades", "Hearts", "Clubs", "Diamonds"}

func (cardStruct *card) getCard() string {
	return cardStruct.face + " of " + cardStruct.suit
}

// GetRandomCard returns a random card as a string.
func GetRandomCard() string {
	randomCard := &card{face: faces[rand.Intn(len(faces))], suit: suits[rand.Intn(len(suits))]}
	return randomCard.getCard()
}

// GetRandomCards returns a slice containing n non-unique cards as strings.
// For a unique set, use getDeck()
// Panics if n <= 0.
func GetRandomCards(n int) []string {
	if n <= 0 {
		panic("Invalid number of cards")
	}
	randomCards := []string{}
	for len(randomCards) < n {
		randomCards = append(randomCards, GetRandomCard())
	}
	return randomCards
}
