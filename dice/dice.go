package dice

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Roll returns random results for num dice of sides faces each
// panics if either argument is <= 0
func Roll(num, sides int) (rolls []int) {
	rolls = []int{}
	if num <= 0 {
		// FIXME String formatting
		panic("Cannot roll 0 dice.")
	}
	if sides <= 2 {
		// FIXME String formatting
		panicMsg := "Cannot roll dice with 0 faces"
		panic(panicMsg)
	}
	for len(rolls) < num {
		rolls = append(rolls, rand.Intn(sides)+1)
	}
	return
}

// RollUneven returns results for a number of dice of sides faces
func RollUneven(sides ...int) (unevenRolls []int, err string) {
	unevenRolls = []int{}
	err = ""
	for i, v := range sides {
		if v > 2 {
			unevenRolls = append(unevenRolls, rand.Intn(v)+1)
			continue
		}
		// FIXME String formatting
		err = err + "Error on input."
		fmt.Printf("Error in input %v.", i)
		unevenRolls = append(unevenRolls, 0)
	}
	return
}
