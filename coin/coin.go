package coin

import (
	"github.com/l3njo/play/misc"
)

// Flip returns a boolean value, true for heads and false for tails.
func Flip() bool {
	return misc.Bool()
}

// FlipVerbose returns the result ofa flips as a string
func FlipVerbose() string {
	if Flip() {
		return "heads"
	}
	return "tails"
}

// Flipn returns num boolean values, true for heads and false for tails.
func Flipn(num int) (resultSet []bool) {
	resultSet = []bool{}
	for len(resultSet) < num {
		resultSet = append(resultSet, misc.Bool())
	}
	return
}

// FlipVerbosen returns the result of num flips as a slice of strings
func FlipVerbosen(num int) (resultSet []string) {
	resultSet = []string{}
	rawSet := Flipn(num)
	for _, isHeads := range rawSet {
		if isHeads {
			resultSet = append(resultSet, "heads")
		} else {
			resultSet = append(resultSet, "tails")
		}
	}
	return
}
