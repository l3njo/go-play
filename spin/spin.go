package spin

import "github.com/l3njo/play/misc"

// Spin returns a random int in the range 1, n
func Spin(n int) int {
	return misc.Int(n) + 1
}

// Verbose returns a random string from a provided sequence or slice of strings
func Verbose(s ...string) string {
	return s[misc.Int(len(s))]
}
