package challenges

import (
	"bufio"
	"os"
	"strings"
)

// D2C1 - Day 2 Challenge 1
// Input is an EOF'd list of strings
// Output is a single unsigned integer
func D2C1() {
	var twice, thrice int

	// This will go on until we reach EOF
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		id := in.Text()
		var found2, found3 bool

		for _, char := range id {
			count := strings.Count(id, string(char))
			found2 = found2 || count == 2
			found3 = found3 || count == 3

			if found2 && found3 {
				// Exit early when we've already reached both results
				break
			}
		}
		if found2 {
			twice++
		}
		if found3 {
			thrice++
		}
	}

	println(twice * thrice)
}

// D2C2 - Day 2 Challenge 2
// Input is an EOF'd list of 26-chars strings
// Output is a single string
func D2C2() {
	// Allocate 250 strings
	ids := make([]string, 250)
	length := 0
	for in := bufio.NewScanner(os.Stdin); in.Scan(); length++ {
		// Populate ids
		ids[length] = in.Text()
	}
	// Shrink to fit
	ids = ids[0:length]

	for i, a := range ids[0 : length/2] {
		for j, b := range ids {
			if j == i {
				// Skip a == b
				continue
			} else if ok, char := diff(a, b); ok {
				println(a[:char] + a[char+1:])
			}
		}
	}

}

// Checks that only one char differs between strings a and b
// Returns (true, pos) where pos is the index in the a string of the differing
// char
func diff(a, b string) (bool, int) {
	cost := 0
	var diff int

	for i, char := range a {
		if rune(b[i]) != char {
			diff = i
			cost++
			if cost > 1 {
				// More than one char, exit early
				return false, 0
			}
		}
	}

	return true, diff
}
