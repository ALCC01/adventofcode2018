package challenges

import (
	"bufio"
	"os"
	"regexp"
	"sort"
	"strconv"
)

var (
	d4regex  = regexp.MustCompile("^.+#([0-9]+).+$")
	d4regex2 = regexp.MustCompile("^.+:([0-9]+).+$")
)

type guard struct {
	Duration int
	Cycles   [][2]int
	Minutes  [60]int
}

// D4C1 - Day 4 Challenge 1
func D4C1() {
	accumulator := []string{}
	line := ""
	// This will go on until we reach EOF
	for in := bufio.NewScanner(os.Stdin); in.Scan(); accumulator = append(accumulator, line) {
		line = in.Text()
	}

	sort.Strings(accumulator)

	guards := map[int]guard{}
	currentGuard, currentCycle := 0, 0
	for _, s := range accumulator {
		switch s[19] {
		case 71: // G
			// New guard
			g, _ := strconv.ParseInt(d4regex.FindStringSubmatch(s)[1], 10, 0)
			currentGuard = int(g)
		case 102: // f
			// Fell asleep
			m, _ := strconv.ParseInt(d4regex2.FindStringSubmatch(s)[1], 10, 0)
			currentCycle = int(m)
		case 119: // w
			// Woke up
			m, _ := strconv.ParseInt(d4regex2.FindStringSubmatch(s)[1], 10, 0)

			duration := int(m) - currentCycle
			cycle := [2]int{currentCycle, int(m)}

			guard := guards[currentGuard]
			guard.Duration += duration
			guard.Cycles = append(guard.Cycles, cycle)
			guards[currentGuard] = guard
		}
	}

	max, maxv := 0, 0
	for k, v := range guards {
		if v.Duration > maxv {
			maxv = v.Duration
			max = k
		}
	}

	println(max * popularMinute(guards[max].Cycles))
}

// D4C2 - Day 4 Challenge 2
func D4C2() {
	accumulator := []string{}
	line := ""
	// This will go on until we reach EOF
	for in := bufio.NewScanner(os.Stdin); in.Scan(); accumulator = append(accumulator, line) {
		line = in.Text()
	}

	sort.Strings(accumulator)

	guards := map[int]guard{}
	currentGuard, currentCycle := 0, 0
	for _, s := range accumulator {
		switch s[19] {
		case 71: // G
			// New guard
			g, _ := strconv.ParseInt(d4regex.FindStringSubmatch(s)[1], 10, 0)
			currentGuard = int(g)
		case 102: // f
			// Fell asleep
			m, _ := strconv.ParseInt(d4regex2.FindStringSubmatch(s)[1], 10, 0)
			currentCycle = int(m)
		case 119: // w
			// Woke up
			m, _ := strconv.ParseInt(d4regex2.FindStringSubmatch(s)[1], 10, 0)

			guard := guards[currentGuard]
			for i := currentCycle; i < int(m); i++ {
				// Vote for the minutes during which the guard has beel asleep
				guard.Minutes[i]++
			}
			guards[currentGuard] = guard
		}
	}

	max, maxm, maxi := 0, 0, 0
	for k, v := range guards {
		i, m := maxMinute(v.Minutes)
		if m > maxm {
			max = k
			maxm = m
			maxi = i
		}
	}

	println(max * maxi)
}

func maxMinute(minutes [60]int) (index, max int) {
	for i, m := range minutes {
		if m > max {
			index = i
			max = m
		}
	}

	return
}

func popularMinute(cycles [][2]int) int {
	votes := [60]int{}
	for _, cycle := range cycles {
		for i := cycle[0]; i < cycle[1]; i++ {
			votes[i]++
		}
	}

	max, maxv := 0, 0
	for k, v := range votes {
		if v > maxv {
			maxv = v
			max = k
		}
	}

	return max
}
