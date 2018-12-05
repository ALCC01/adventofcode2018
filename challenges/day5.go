package challenges

import (
	"fmt"
	"strings"
	"unicode"
)

// D5C1 - Day 5 Challenge 1
// Input is a single string
// Output is a single unsigned int
func D5C1() {
	var new string
	fmt.Scanln(&new)

	println(len(react(new, -1)))
}

// D5C2 - Day 5 Challenge 2
// Input is a single string
// Output is a single unsigned int
func D5C2() {
	var new string
	fmt.Scanln(&new)
	min := len(new)
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for _, a := range alphabet {
		l := len(react(new, a))

		if l < min {
			min = l
		}
	}

	println(min)
}

func react(new string, filter rune) (old string) {
	// Continue until we don't have any other edits to make
	for old != new {
		old = new
		dropnext := false
		i := 0
		new = strings.Map(func(r rune) rune {
			i++
			// We don't have another rune to compare the current one to
			if i == len(old) {
				return r
			}
			// Drop the current rune
			if dropnext {
				dropnext = false
				return -1
			}
			// Filter out the filter rune
			if unicode.ToLower(r) == filter {
				return -1
			}

			next := rune(old[i])
			opposite := unicode.IsLower(r) != unicode.IsLower(next) // XOR
			same := unicode.ToLower(r) == unicode.ToLower(next)

			// Drop the runes if they are both of opposite casing AND are the
			// same letter
			if opposite && same {
				dropnext = true
				return -1
			}

			// Keep the rune
			return r
		}, old)
	}

	return
}
