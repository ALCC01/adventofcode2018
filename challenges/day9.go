package challenges

import (
	"container/ring"
	"fmt"
)

// D9C1 - Day 9 Challenge 1
// Input is a string in the form of `464 players; last marble is worth 70918 points`
// Output is a single unsigned integer
func D9C1() {
	var players, marbles int
	fmt.Scanf("%d players; last marble is worth %d points", &players, &marbles)

	println(playCircle(players, marbles))
}

// D9C2 - Day 9 Challenge 2
// Input is a string in the form of `464 players; last marble is worth 70918 points`
// Output is a single unsigned integer
func D9C2() {
	var players, marbles int
	fmt.Scanf("%d players; last marble is worth %d points", &players, &marbles)

	println(playCircle(players, marbles*100))
}

func playCircle(players, marbles int) int {
	// Ring is a circular list
	circle := ring.New(1)
	// Central marble has value 0
	circle.Value = 0
	scores := make([]int, players)

	for next := 1; next <= marbles; next++ {
		// If the marble to be inserted is a multiple of 23
		if next%23 == 0 {
			// Move to the 7th coutner-clockwise marble
			circle = circle.Move(-8)
			// Remove it
			scores[next%players] += next + circle.Unlink(1).Value.(int)
			// Move to the new central marble
			circle = circle.Move(1)
		} else {
			new := ring.New(1)
			new.Value = next

			// Move to the insertion point
			circle = circle.Move(1)
			circle.Link(new)
			// Move to the new central point
			circle = circle.Move(1)
		}
	}

	max := 0
	for _, v := range scores {
		if v > max {
			max = v
		}
	}

	return max
}
