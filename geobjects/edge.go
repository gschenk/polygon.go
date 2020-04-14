// Package geobjects provides geometric objects, viz. nodes, edges, and polygon
package geobjects

import "polyGo/vector"

// Edge provides the position vector of one point and the edge vector
type Edge struct {
	node   Node
	vector vector.Vec
}

// MakeEdge returns an Edge struct when provided with position vector of two
// points A and B. The Edge vector is oriented from a to b.
func MakeEdge(a, b vector.Vec) Edge {
	return Edge{MakeNode(a), vector.FromAtoB(a, b)}
}
