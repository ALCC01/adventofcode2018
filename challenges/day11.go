package challenges

import (
	"fmt"
)

// D11C1 - Day 11 Challenge 1
// Input is a single unsigned int
// Output is a single string
func D11C1() {
	serial := 0
	fmt.Scanf("%d", &serial)
	sum := makeSum(serial)
	X, Y, _ := maxArea(sum, 3, 3)
	fmt.Printf("%d,%d\n", X, Y)
}

// D11C2 - Day 11 Challenge 2
// Input is a single unsigned int
// Output is a single string
func D11C2() {
	serial := 0
	fmt.Scanf("%d", &serial)
	sum := makeSum(serial)
	X, Y, S := maxArea(sum, 3, 300)
	fmt.Printf("%d,%d,%d\n", X, Y, S)
}

// The idea is to have each cell of this grid have as a value the sum of all
// the cells that are below it and to the left, so that a bigger area can
// be mapped to a single cell. The value of an area [(x, y), (X, Y)] should be
// easy to calculate as it is basically a big sum of the values of the two
// corners minus the areas that are either lower or lefterer than (x, y).
func makeSum(serial int) (sum [301][301]int) {
	for y := 1; y <= 300; y++ {
		for x := 1; x <= 300; x++ {
			i := (((((x+10)*y + serial) * (x + 10)) / 100) % 10) - 5
			// The value of this thing, plus that of the cells below it
			// and those to the left of it, minus their intersection
			sum[y][x] = i + sum[y-1][x] + sum[y][x-1] - sum[y-1][x-1]
		}
	}

	return
}

func maxArea(sum [301][301]int, minRange, maxRange int) (X, Y, S int) {
	max := 0
	for s := minRange; s <= maxRange; s++ {
		for y := 0; y <= 299-s; y++ {
			for x := 0; x <= 299-s; x++ {
				// The value of this corner, plus that of the opposite one
				// minus the cells we don't care about
				m := sum[y][x] + sum[y+s][x+s] - sum[y+s][x] - sum[y][x+s]
				if m > max {
					S, X, Y, max = s, x+1, y+1, m
				}
			}
		}
	}
	println(max)

	return
}
