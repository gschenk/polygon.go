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
	Base       Node
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
	det := vector.SafeDet(evec, toCent)
	return Edge{
		id:         id(),
		pos:        a.v,
		evec:       evec,
		Centre:     cent,
		det:        det,
		sign:       math.Signbit(det),
		Base:       a,   // node at base
		Head:       b,   // node at head
		Next:       b.v, // defaults to position of head, edge
		nextAngle:  0,
		HasOutside: false, // set this to true when there is a point Outside this edge
		Outside:    vector.Vecs{},
	}
}

// Warning: Method modifies Edge struct!
// processPoint determines if a point is outside or inside of the edge
// for outside points it also checks if it is the next node
func (e *Edge) processPoint(p vector.Vec) (vector.Vec, bool) {

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

// FindOutsidePoints runs the outside point search for a slice of points
func (e *Edge) FindOutsidePoints(ps vector.Vecs) vector.Vecs {
	insidePs := vector.Vecs{}
	for _, p := range ps {
		insideP, isInside := e.processPoint(p)
		if isInside {
			insidePs = append(insidePs, insideP)
		}
	}
	return insidePs // next search only through points centre-wards of the present edge
}

// Edges Methods

// dets method returns a slice of det values
func (es Edges) dets() []float64 {
	ys := make([]float64, len(es))
	for i, e := range es {
		ys[i] = e.det
	}
	return ys
}

// signs method returns a slice of bool values, `true` for negative det value
func (es Edges) signs() []bool {
	bs := make([]bool, len(es))
	for i, e := range es {
		bs[i] = e.sign
	}
	return bs
}
