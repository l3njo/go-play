package dice

import (
	"fmt"

	"github.com/l3njo/play/misc"
)

// Roll returns random results for num dice of sides faces each
// Returns an error if num <= 0 || sides <= 2
func Roll(num, sides int) (rolls []int, err string) {
	rolls = []int{}
	err = ""
	if num <= 0 {
		err = err + fmt.Sprintf("Cannot roll %v dice.\n", num)
		return
	}
	if sides <= 2 {
		err = err + fmt.Sprintf("Cannot roll dice with %v faces.\n", sides)
		return
	}
	for len(rolls) < num {
		rolls = append(rolls, misc.Int(sides)+1)
	}
	return
}

// RollUneven returns results for a number of dice of sides faces
// Returns 0 and an error for any side that is <= 2
func RollUneven(sides ...int) (unevenRolls []int, err string) {
	unevenRolls = []int{}
	err = ""
	for i, v := range sides {
		if v > 2 {
			unevenRolls = append(unevenRolls, misc.Int(v)+1)
			continue
		}
		err = err + fmt.Sprintf("Error on input %v.\n", i)
		unevenRolls = append(unevenRolls, 0)
	}
	return
}
