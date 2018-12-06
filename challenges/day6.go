package challenges

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// D6C1 - Day 6 Challenge 1
// Input is a list of strings in the form of `346, 260`
// Output is a single unsigned int
func D6C1() {
	var coords [][2]int
	maxX, maxY, minX, minY := 0, 0, 0, 0

	// Populate coords
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		x, y := 0, 0
		fmt.Sscanf(in.Text(), "%d, %d", &x, &y)
		coords = append(coords, [2]int{x, y})

		// Register max and min coords
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		if x < minX {
			minX = x
		}
		if y < minY {
			minY = y
		}
	}

	votes := map[int]int{}
	// Calculate the area of each point
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			c := findClosest([2]int{x, y}, coords)

			if y == minY || y == maxY || x == maxX || x == minX {
				// Weight down points that have infinite area
				votes[c] = -1e4
			} else if c == -1 {
				// Ignore draws
				continue
			} else {
				votes[c]++
			}
		}
	}

	// Find the largest area
	max := -1
	for _, v := range votes {
		if v > max {
			max = v
		}
	}

	println(max)
}

// D6C2 - Day 6 Challenge 2
// Input is a list of strings in the form of `346, 260`
// Output is a single unsigned int
func D6C2() {
	var coords [][2]int
	maxD := int(1e4)
	maxX, maxY, minX, minY := 0, 0, 0, 0

	// Populate coords
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		x, y := 0, 0
		fmt.Sscanf(in.Text(), "%d, %d", &x, &y)
		coords = append(coords, [2]int{x, y})

		// Register max and min coords
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		if x < minX {
			minX = x
		}
		if y < minY {
			minY = y
		}
	}

	votes := 0
	// This assumes that the points are far enough from each other that we don't
	// need to look outside of the [(minX,minY), (maxX,maxY)] rectangles.
	// It should probably include a buffer zone calculated using maxD and the
	// distance between the points that are the farthest from each other on the
	// X or Y axis,
	// Works for me, so I'll spare my CPU the extra loops
	for y := minY; y <= maxY; y++ {
	outer:
		for x := minX; x <= maxX; x++ {
			accumulator := 0
			for _, c := range coords {
				accumulator += manhattanDistance(c, [2]int{x, y})
				if accumulator >= maxD {
					// We don't need to continue calculating
					continue outer
				}
			}
			// This point survived
			votes++
		}
	}

	println(votes)
}

func findClosest(p [2]int, coords [][2]int) int {
	minDistance, closest := math.MaxInt64, -1

	for n, p0 := range coords {
		if p0 == p {
			// Exit early if it is the same point
			return n
		}
		d := manhattanDistance(p, p0)

		if d < minDistance {
			// Found a closer point
			minDistance = d
			closest = n
		} else if d == minDistance {
			// If we don't find another point that's closer, this is a draw
			closest = -1
		}
	}

	return closest
}

// |X1-X2| + |Y1-Y2|
func manhattanDistance(p1, p2 [2]int) int {
	return int(math.Abs(float64(p1[0]-p2[0])) + math.Abs(float64(p1[1]-p2[1])))
}
