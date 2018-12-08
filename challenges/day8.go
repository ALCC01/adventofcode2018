package challenges

import (
	"fmt"
)

// D8C1 - Day 8 Challenge 1
// Input is a list of integers
// Output is a single unsigned integer
func D8C1() {
	println(digestNodes(1))
}

// D8C2 - Day 8 Challenge 2
// Input is a list of integers
// Output is a single unsigned integer
func D8C2() {
	println(createTree(1)[0].Value)
}

type node struct {
	Child []node
	Meta  []int
	Value int
}

// Recursively calculate the value of each node
func createTree(n int) (nodes []node) {
	for i := 0; i < n; i++ {
		header := digest(2)
		n := node{
			createTree(header[0]),
			digest(header[1]),
			0,
		}
		n.Value = value(n)
		nodes = append(nodes, n)
	}

	return
}

func value(n node) (value int) {
	// If the node has no children, its value is the sum of its metadata values
	if len(n.Child) == 0 {
		return sum(n.Meta)
	}
	// Otherwise, it is the sum of the values of its child nodes
	for _, i := range n.Meta {
		if i <= len(n.Child) {
			value = value + n.Child[i-1].Value
		}
	}

	return
}

// Recursively digests n nodes, returning the sum of their metadata values
func digestNodes(n int) (meta int) {
	for i := 0; i < n; i++ {
		header := digest(2)
		meta = meta + digestNodes(header[0])
		meta = meta + sum(digest(header[1]))
	}

	return
}

// Ingests n ints from stdin
func digest(n int) []int {
	accumulator := []int{}
	for i := 0; i < n; i++ {
		a := 0
		fmt.Scan(&a)
		accumulator = append(accumulator, a)
	}
	return accumulator
}

func sum(a []int) (r int) {
	for i := 0; i < len(a); i++ {
		r = r + a[i]
	}
	return
}
