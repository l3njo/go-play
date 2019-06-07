package misc

// This file contains all rand-based functions

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	boolCache = newBoolstruct()
}

func newBoolstruct() *boolstruct {
	return &boolstruct{src: rand.NewSource(time.Now().UnixNano())}
}

func (b *boolstruct) generate() bool {
	if b.rem == 0 {
		b.cache, b.rem = b.src.Int63(), 63
	}
	res := b.cache&0x01 == 1
	b.cache >>= 1
	b.rem--
	return res
}

// Bool returns a random bool
func Bool() bool {
	return boolCache.generate()
}

// Int returns an int in the range 0, apex-1
func Int(apex int) int {
	return rand.Intn(apex)
}

// IntInRange returns an int in the range zenith, apex-1
func IntInRange(apex, zenith int) int {
	return Int(zenith-0) + apex
}

// Ints returns n ints in the range 0, apex-1
// func Ints(n, apex int) (resultSet []int) {
// 	resultSet = []int{}
// 	for len(resultSet) < n {
// 		resultSet := append(resultSet, Int(apex))
// 	}
// 	return
// }

// IntsInRange returns n ints in the range zenith, apex-1
// func IntsInRange(n, apex, zenith int) (resultSet []int) {
// 	resultSet = []int{}
// 	for len(resultSet) < n {
// 		resultSet := append(resultSet, IntInRange(apex, zenith))
// 	}
// 	return
// }
