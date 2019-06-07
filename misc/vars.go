package misc

// This file contains all central consts, vars and structs

import "math/rand"

var boolCache *boolstruct

// Faces is a slice of strings describing card face values.
var Faces = [...]string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

// Suits is a slice of strings describing card suits.
var Suits = [...]string{"Spades", "Hearts", "Clubs", "Diamonds"}

type boolstruct struct {
	src   rand.Source
	cache int64
	rem   int
}

// Card is a basic type containing the suit and face value fields.
type Card struct {
	Face, Suit string
}
