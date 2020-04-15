// Package geo provides geometric objects, viz. nodes, edges, and polygon
package geo

import (
	"fmt"
	"polyGo/vector"
)

// Polygon is a struct that collects nodes, edges, and centre of a polygon
type Polygon struct {
	id       int64
	Nodes    []Node
	Edges    []Edge
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

// linkEdges returns edges for a set of nodes
func linkEdges(as []Node) []Edge {

	// reordered as, first element last
	bs := append(as[1:], as[0])

	es := make([]Edge, len(as))
	for i := range as {
		es[i] = NewEdge(as[i], bs[i])
	}
	return es
}

func checkCompleteness(es []Edge) bool {
	fmt.Println(es[0].tail)
	fmt.Println(es[len(es)-1].head)
	return es[0].tail == es[len(es)-1].head
}

// NewPoly forms a polygon struct from a unique set of Nodes
func NewPoly(nodes []Node) Polygon {
	edges := linkEdges(nodes)
	return Polygon{
		id:       id(),
		Nodes:    nodes,
		Edges:    edges,
		Centre:   findCentre(nodes),
		Complete: checkCompleteness(edges),
	} //stub
}
