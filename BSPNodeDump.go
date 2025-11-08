package main

import "fmt"

type BSPNodeDump struct {
	// Rather that knowing who
	// this node's children is,
	// we know the parent
	parent int

	// Indicies of the segments that
	// work in this node
	lines []int
}

func (node *BSPNodeDump) String() string {
	out := fmt.Sprintf("NODE %d", node.parent)

	for i := range len(node.lines) {
		out += fmt.Sprintf(" %d", node.lines[i])
	}

	return out
}
