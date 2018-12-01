package challenges

import (
	"bufio"
	"os"
	"strconv"
)

// D1C1 - Day 1 Challenge 1
// Input is an EOF'd list of signed integers
// Output is a single signed integer
func D1C1() {
	var accumulator, delta int64

	// This will go on until we reach EOF
	for in := bufio.NewScanner(os.Stdin); in.Scan(); accumulator += delta {
		// Assume the input is properly formatted
		delta, _ = strconv.ParseInt(in.Text(), 10, 64)
	}

	println(accumulator)
}

// D1C2 - Day 1 Challenge 2
// Input is an EOF'd list of signed integers
// Output is a single signed integer
func D1C2() {
	var accumulator, length int
	// Allocate 1200 deltas, hopefully they're enough
	deltas := make([]int, 1200)
	frequencies := make(map[int]bool)
	frequencies[0] = true

	for in := bufio.NewScanner(os.Stdin); in.Scan(); length++ {
		// Assume the input is properly formatted
		delta, _ := strconv.ParseInt(in.Text(), 10, 0)
		accumulator += int(delta)
		deltas[length] = int(delta)

		if frequencies[accumulator] {
			println(accumulator)
			os.Exit(0)
		} else {
			frequencies[accumulator] = true
		}
	}

	// Shrink to fit delta's elements
	deltas = deltas[0:length]
	for i := 0; ; i++ {
		// We could loop a few times through the array
		delta := deltas[i%len(deltas)]
		accumulator += delta

		if frequencies[accumulator] {
			println(accumulator)
			os.Exit(0)
		} else {
			frequencies[accumulator] = true
		}
	}
}
