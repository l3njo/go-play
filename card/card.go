package card

import (
	"math/rand"
)

type card struct {
	face, suit string
}

var faces = [...]string{""}
var suits = [...]string{""}

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
