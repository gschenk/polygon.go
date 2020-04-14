// Package geobjects provides geometric objects, viz. nodes, edges, and polygon
package geobjects

import (
	"polyGo/vector"
)

// Polygon is a struct that collects nodes, edges, and centre of a polygon
type Polygon struct {
	Id       int64
	Nodes    []Node
	Centre   vector.Vec
	Complete bool
}

// findCentre returns the geometric centre of a set of nodes
func findCentre(ns []Node) vector.Vec {

	// get position vectors from nodes
	vs := mapVecFunToNodes(
		func(n Node) vector.Vec { return n.v },
		ns,
	)

	// sum of all vectors
	vSum := vector.FoldrVecs(vector.Sum, vs, vector.Zero)

	// normalise
	s := 1 / float64(len(ns))
	return vector.ScalarMult(s, vSum)
}

// NewPoly forms a polygon struct from a unique set of Nodes
func NewPoly(ns []Node) Polygon {
	return Polygon{
		Id:       id(),
		Nodes:    ns,
		Centre:   findCentre(ns),
		Complete: false,
	} //stub
}
