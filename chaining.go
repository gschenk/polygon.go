package main

import (
	"polyGo/geo"
)

// Note: this is a recursive function
// find outer edges of the complex hull with points on one side of a give edge
// initial call with base node of parent edge in nodes
func recChainSearch(nodes geo.Nodes, edge geo.Edge, depth uint) (geo.Nodes, uint) {

	if depth >= maxrecursion {
		return nodes, depth
	}

	// when an edge has no outside points, we are done here
	if !edge.HasOutside {
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

	// call FindOutsidePoints() on testEdge to store its outside points
	testEdge.FindOutsidePoints(edge.Outside)

	depth += 1
	return recChainSearch(moreNodes, testEdge, depth)
}

func chainSearch(edges geo.Edges) geo.Nodes {
	result := geo.Nodes{}
	for _, edge := range edges {
		nextNodes, _ := recChainSearch(
			geo.Nodes{edge.Base}, // slice of 1 Node
			edge,
			0,
		)
		result = append(result, nextNodes...)
	}
	return result
}
