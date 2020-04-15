// Package geo provides geometric objects, viz. nodes, edges, and polygon
package geo

import (
	"math"
	"polyGo/tools"
	"polyGo/vector"
)

// Edge provides the position vector of one point and the edge vector
type Edge struct {
	id         int64
	pos        vector.Vec // position of base node
	evec       vector.Vec // edge vector, base to head
	det        float64
	sign       bool // true for a negative det
	base       int64
	head       int64
	next       vector.Vec // position of next outside node, defaults to head
	nextAngle  float64
	nextExists bool
	outside    vector.Vecs
}

// Edges provide a list of edges
type Edges []Edge

// Edge Methods

// NewEdge returns an Edge struct when provided with position vector of two
// points A and B. The Edge vector is oriented from a to b.
func NewEdge(a, b Node, cent vector.Vec) Edge {
	evec := vector.FromAtoB(a.v, b.v)
	toCent := vector.FromAtoB(a.v, cent)
	det := vector.Det(evec, toCent)
	return Edge{
		id:         id(),
		pos:        a.v,
		evec:       evec,
		det:        det,
		sign:       math.Signbit(det),
		base:       a.id, // id of node at base
		head:       b.id, // id of node at head
		next:       b.v,  // defaults to position of head, edge
		nextAngle:  0,
		nextExists: false, // set this to true when there is a point outside this edge
		outside:    vector.Vecs{},
	}
}

// Warning: Method modifies Edge struct!
// findOutsidePoints
func (e *Edge) findOutsidePoints(p vector.Vec) (vector.Vec, bool) {

	// vector vp from edge (base) to point p
	vp := vector.FromAtoB(e.pos, p)

	// find determinant of edge vector and vp
	det := vector.Det(e.evec, vp)

	// test if point is on edge, inside, or outside
	if tools.FloatIsZero(det) {
		// colinear vector, ignore point
		// (With Akl-Toussaint heuristic it has to be on the edge)
		return vector.Zero, false
	} else if e.sign == (det < 0) {
		// inside point, return it, for next edge to check
		return p, true
	}
	// outside, don't return it, store it with edge
	e.outside = append(e.outside, p)

	// and check if it is next node

	angle := math.Pow(det, 2) / vp.Norm2()
	// this is not _really_ the angle
	// when restricting to [0 .. pi] it is bijective mapping to it

	if angle > e.nextAngle {
		e.next = p
		e.nextAngle = angle
		e.nextExists = true
	}

	return vector.Zero, false
}

// Edges Methods

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

// Points method returns a slice of vectors for each node,
func (es Edges) Points() []vector.Vecs {
	return mapVecsFunToEdges(
		func(e Edge) vector.Vecs { return e.outside },
		es,
	)
}
