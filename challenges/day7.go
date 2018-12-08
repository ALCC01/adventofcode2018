package challenges

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// D7C1 - Day 7 Challenge 1
// Input is a list of strings in the form of `Step I must be finished before step Q can begin.`
// Output is a single string
func D7C1() {
	mustWait := map[rune]string{}

	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		var before, after rune
		fmt.Sscanf(in.Text(), "Step %c must be finished before step %c can begin.\n", &before, &after)
		// Update the list of steps `after` needs to have been finished before
		// it can be started
		mustWait[after] = mustWait[after] + string(before)
		// Explicitly set nil values so that we can iterate over them later
		if mustWait[before] == "" {
			mustWait[before] = ""
		}
	}

	old, new := "a", ""
	for old != new {
		old = new
		candidates := []rune{}
		for id, waits := range mustWait {
			if waits == "!" {
				// Step has already been completed
				continue
			}
			if finished(id, mustWait) {
				// Add this step to the list of candidates
				candidates = append(candidates, id)
			}
		}
		// If we still have candidates
		if len(candidates) != 0 {
			// Set the first in alphabetical order as completed
			sort.Slice(candidates, func(i, j int) bool {
				return candidates[i] < candidates[j]
			})
			finished := candidates[0]
			mustWait[finished] = "!"
			new = new + string(finished)
		}
	}
	println(old)
}

func finished(id rune, mustWait map[rune]string) bool {
	if mustWait[id] == "" {
		return true
	}
	for i := 0; i < len(mustWait[id]); i++ {
		if mustWait[rune(mustWait[id][i])] != "!" {
			return false
		}
	}

	return true
}
