package card

import (
	"math/rand"
	"time"
)

// Card is a basic type containing the suit and face value fields.
type Card struct {
	face, suit string
}

// TOD0 make constants
var faces = [...]string{"Ace", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
var suits = [...]string{"Spades", "Hearts", "Clubs", "Diamonds"}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// GetCard ...
func (cardStruct *Card) GetCard() string {
	return cardStruct.face + " of " + cardStruct.suit
}

// GetRandomCard returns a random card as a string.
func GetRandomCard() string {
	randomCard := &Card{face: faces[rand.Intn(len(faces))], suit: suits[rand.Intn(len(suits))]}
	return randomCard.GetCard()
}

// GetRandomCards returns a slice containing n non-unique cards as strings.
// For a unique set, use GetDeck()
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
