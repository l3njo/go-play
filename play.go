package main

import (
	"fmt"

	"github.com/l3njo/play/dice"
)

func main() {
	fmt.Println(dice.Roll(2, 3))
	fmt.Println(dice.RollUneven(2, 3, 4, 2, 7, 3))
}
