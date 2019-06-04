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

// GetRandomCard returns a random card in string form
func GetRandomCard() string {
	newCard := new(card)
	newCard.face = faces[rand.Intn(len(faces))]
	newCard.suit = suits[rand.Intn(len(suits))]
	return newCard.getCard()
}
