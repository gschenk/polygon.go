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
	Centre     vector.Vec
	det        float64
	sign       bool // true for a negative det
	base       Node
	Head       Node
	Next       vector.Vec // position of Next Outside node, defaults to head
	nextAngle  float64
	HasOutside bool
	Outside    vector.Vecs
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
		Centre:     cent,
		det:        det,
		sign:       math.Signbit(det),
		base:       a,   // node at base
		Head:       b,   // node at head
		Next:       b.v, // defaults to position of head, edge
		nextAngle:  0,
		HasOutside: false, // set this to true when there is a point Outside this edge
		Outside:    vector.Vecs{},
	}
}

// Warning: Method modifies Edge struct!
// FindOutsidePoints
func (e *Edge) FindOutsidePoints(p vector.Vec) (vector.Vec, bool) {

	// vector vp from edge (base) to point p
	vp := vector.FromAtoB(e.pos, p)

	// find determinant of edge vector and vp
	det := vector.Det(e.evec, vp)

	// test if point is on edge, inside, or Outside
	if tools.FloatIsZero(det) {
		// colinear vector, ignore point
		// (With Akl-Toussaint heuristic it has to be on the edge)
		return vector.Zero, false
	} else if e.sign == (det < 0) {
		// inside point, return it, for Next edge to check
		return p, true
	}
	// Outside, don't return it, store it with edge
	e.Outside = append(e.Outside, p)

	// and check if it is Next node

	angle := math.Pow(det, 2) / vp.Norm2()
	// this is not _really_ the angle
	// when restricting to [0 .. pi] it is bijective mapping to it

	if angle > e.nextAngle {
		e.Next = p
		e.nextAngle = angle
		e.HasOutside = true
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
		func(e Edge) vector.Vecs { return e.Outside },
		es,
	)
}
