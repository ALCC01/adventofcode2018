package challenges

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

var d3regex = regexp.MustCompile(`^#[0-9]+ @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)$`)

// D3C1 - Day 3 Challenge 1
// Input is a list of strings in the form of `#1 @ 469,741: 22x26`
// Output is a single unsigned integer
func D3C1() {
	var (
		m       [2000][2000]int64
		counter int
	)

	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		x, y, w, h := parseClaim(in.Text())

		// Apply all claims to the matrix
		for i := y; i < y+h; i++ {
			for j := x; j < x+w; j++ {
				switch m[i][j] {
				case 2:
					break
				case 1:
					// This tile reached 2 overlapping claims, increase counter
					counter++
					fallthrough
				case 0:
					m[i][j]++
				}
			}
		}
	}

	println(counter)
}

// D3C2 - Day 3 Challenge 2
// Input is a list of strings in the form of `#1 @ 469,741: 22x26`
// Output is a single unsigned integer
func D3C2() {
	claims := map[int][4]int64{}
	var (
		m [2000][2000]int64
		n int
	)

	for in := bufio.NewScanner(os.Stdin); in.Scan(); n++ {
		x, y, w, h := parseClaim(in.Text())
		untouched := true

		// Apply all claims to the matrix, keep only those that do not already
		// overlap with another one
		for i := y; i < y+h; i++ {
			for j := x; j < x+w; j++ {
				switch m[i][j] {
				default:
					untouched = false
					break
				case 1:
					untouched = false
					fallthrough
				case 0:
					m[i][j]++
				}
			}
		}

		if untouched {
			claims[n+1] = [4]int64{
				x, x + w,
				y, y + h,
			}
		}
	}

outer:
	for k, v := range claims {
		x, X, y, Y := v[0], v[1], v[2], v[3]

		for i := y; i < Y; i++ {
			for j := x; j < X; j++ {
				if m[i][j] != 1 {
					// More than one claims on this tile
					continue outer
				}
			}
		}

		// Found a claim with all tiles == 1
		println(k)
	}
}

func parseClaim(claim string) (x, y, w, h int64) {
	fragments := d3regex.FindStringSubmatch(claim)
	x, _ = strconv.ParseInt(fragments[1], 10, 0)
	y, _ = strconv.ParseInt(fragments[2], 10, 0)
	w, _ = strconv.ParseInt(fragments[3], 10, 0)
	h, _ = strconv.ParseInt(fragments[4], 10, 0)

	return
}
