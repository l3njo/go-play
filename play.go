package main

import (
	"fmt"

	"github.com/l3njo/play/card"
)

func main() {
	fmt.Println(card.GetRandomCard())
	for i, v := range card.GetShuffledDeck() {
		fmt.Printf("%v: %v\n", i+1, v)
	}
}
