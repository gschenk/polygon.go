// Package geo provides geometric objects, viz. nodes, edges, and polygon
package geo

import (
	"math"
	"polyGo/vector"
)

// Edge provides the position vector of one point and the edge vector
type Edge struct {
	id   int64
	pos  vector.Vec // position of tail node
	evec vector.Vec // edge vector, tail to head
	det  float64
	sign bool // true for a negative det
	tail int64
	head int64
}

// Edges provide a list of edges
type Edges []Edge

// NewEdge returns an Edge struct when provided with position vector of two
// points A and B. The Edge vector is oriented from a to b.
func NewEdge(a, b Node, cent vector.Vec) Edge {
	evec := vector.FromAtoB(a.v, b.v)
	toCent := vector.FromAtoB(a.v, cent)
	det := vector.Det(evec, toCent)
	return Edge{
		id:   id(),
		pos:  a.v,
		evec: evec,
		det:  det,
		sign: math.Signbit(det),
		tail: a.id, // id of node at tail
		head: b.id, // id of node at head
	}
}

// dets method returns a slice of det values
func (es Edges) dets() []float64 {
	return mapFloatFunToEdges(
		func(e Edge) float64 { return e.det },
		es,
	)
}

// signs method returns a slice of bool values, `true` for negative det value
func (es Edges) signs() []bool {
	return mapBoolFunToEdges(
		func(e Edge) bool { return e.sign },
		es,
	)
}
