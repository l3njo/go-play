package main

import (
	"fmt"
	"play/card"
)

func main() {
	fmt.Println(card.GetRandomCard())
	for i, v := range card.GetDeck() {
		fmt.Printf("%v: %v\n", i+1, v)
	}
}
