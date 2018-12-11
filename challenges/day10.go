package challenges

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type star struct {
	p0 complex128
	v  complex128
}

func (s star) position(tick int) complex128 {
	return s.p0 + complex(float64(tick), 0)*s.v
}

// D10C1 - Day 10 Challenge 1
// Input is a list of strings in the form of `position=<%d,%d> velocity=<%d,%d>.`
// Output is a list of strings
func D10C1() {
	stars := make([]star, 350)
	counter := 0

	for in := bufio.NewScanner(os.Stdin); in.Scan(); counter++ {
		var a, b, c, d int
		fmt.Sscanf(in.Text(), "position=<%d,%d> velocity=<%d,%d>.\n", &a, &b, &c, &d)
		stars[counter] = star{
			complex(float64(a), float64(b)),
			complex(float64(c), float64(d)),
		}
	}
	stars = stars[:counter]

	minx := math.MaxFloat64
	minxt := 0
	for i := 0; i < 15000; i++ {
		a := disperison(positions(stars, i))
		if a < minx {
			minx = a
			minxt = i
		}
	}

	drawStars(stars, minxt)
}

// D10C2 - Day 10 Challenge 2
// Input is a list of strings in the form of `position=<%d,%d> velocity=<%d,%d>.`
// Output is a single unsigned integer
func D10C2() {
	stars := make([]star, 350)
	counter := 0

	for in := bufio.NewScanner(os.Stdin); in.Scan(); counter++ {
		var a, b, c, d int
		fmt.Sscanf(in.Text(), "position=<%d,%d> velocity=<%d,%d>.\n", &a, &b, &c, &d)
		stars[counter] = star{
			complex(float64(a), float64(b)),
			complex(float64(c), float64(d)),
		}
	}
	stars = stars[:counter]

	minx := math.MaxFloat64
	minxt := 0
	for i := 0; i < 15000; i++ {
		a := disperison(positions(stars, i))
		if a < minx {
			minx = a
			minxt = i
		}
	}

	println(minxt)
}

func drawStars(stars []star, tick int) {
	m := [500][500]int{}
	p := positions(stars, tick)
	a, b := bounding(p)

	for _, v := range p {
		m[int(imag(v))-b[0]][int(real(v))-a[0]] = 1
	}

	for i, y := range m {
		if i > b[1]-b[0] {
			break
		}
		for j, x := range y {
			if j > a[1]-a[0] {
				break
			}
			if x == 1 {
				print("#")
			} else {
				print(" ")
			}
		}
		println("")
	}
}

func bounding(positions []complex128) ([2]int, [2]int) {
	x, X := math.MaxInt64, 0
	y, Y := math.MaxInt64, 0

	for _, v := range positions {
		a, b := int(real(v)), int(imag(v))
		if a > X {
			X = a
		}
		if a < x {
			x = int(a)
		}
		if b > Y {
			Y = int(b)
		}
		if b < y {
			y = int(b)
		}
	}

	return [2]int{x, X}, [2]int{y, Y}
}

// Standard deviation over x
func disperison(positions []complex128) float64 {
	average := 0 + 0i
	for _, v := range positions {
		average += v
	}
	average = average / complex(float64(len(positions)), 0)
	var X float64
	for _, v := range positions {
		X += math.Pow(real(v)-real(average), 2)
	}

	return math.Sqrt(X / float64(len(positions)))
}

func positions(stars []star, tick int) (r []complex128) {
	r = make([]complex128, len(stars))
	for k, v := range stars {
		r[k] = v.position(tick)
	}

	return
}
