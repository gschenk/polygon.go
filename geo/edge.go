// Package geo provides geometric objects, viz. nodes, edges, and polygon
package geo

import "polyGo/vector"

// Edge provides the position vector of one point and the edge vector
type Edge struct {
	id   int64
	pos  vector.Vec // position of tail node
	vec  vector.Vec // edge vector, tail to head
	tail int64
	head int64
}

// NewEdge returns an Edge struct when provided with position vector of two
// points A and B. The Edge vector is oriented from a to b.
func NewEdge(a, b Node) Edge {
	return Edge{
		id:   id(),
		pos:  a.v,
		vec:  vector.FromAtoB(a.v, b.v),
		tail: a.id, // id of node at tail
		head: b.id, // id of node at head
	}
}
