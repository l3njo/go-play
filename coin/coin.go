package coin

import (
	"math/rand"
	"time"
)

var flipset *set

type set struct {
	src   rand.Source
	cache int64
	rem   int
}

// bool generates 64 bool values, and caches them in a Int63 var.
// values are popped from cache when a random bool is requested.
// values are regenerated when cache is exhausted.
func (s *set) bool() bool {
	if s.rem == 0 {
		s.cache, s.rem = s.src.Int63(), 63
	}
	res := s.cache&0x01 == 1
	s.cache >>= 1
	s.rem--
	return res
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	flipset = newSet()
}

func newSet() *set {
	return &set{src: rand.NewSource(time.Now().UnixNano())}
}

// Flip returns a boolean value, true for heads and false for tails.
func Flip() bool {
	return flipset.bool()
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
		resultSet = append(resultSet, flipset.bool())
	}
	return
}

// FlipVerbosen returns the result of num flips as a slice of strings
func FlipVerbosen(num int) (resultSet []string) {
	resultSet = []string{}
	rawSet := Flipn(num)
	for _, isHeads := range rawSet {
		// FIXME
		if isHeads {
			resultSet = append(resultSet, "heads")
		} else {
			resultSet = append(resultSet, "tails")
		}
	}
	return
}
