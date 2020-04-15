package main

import (
	"polyGo/geo"
)

// Note: this is a recursive function
// find outer edges of the complex hull with points on one side of a give edge
// initial call with base node of parent edge in nodes
func chaining(nodes geo.Nodes, edge geo.Edge, depth uint) (geo.Nodes, uint) {

	// when an edge has no outside points, we are done here
	if edge.HasOutside {
		return nodes, depth
	}

	// otherwise:

	// create a test edge from present edge's `next` point to its `head`
	nextNode := geo.NewNode(edge.Next)
	testEdge := geo.NewEdge(
		nextNode,
		edge.Head,
		edge.Centre,
	)

	// add nextNode to growing slice of nodes
	moreNodes := append(nodes, nextNode)

	// TODO run FindOutsidePoints() on testEdge

	depth += 1
	return chaining(moreNodes, testEdge, depth)
}
