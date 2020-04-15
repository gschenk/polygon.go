// Package geo provides geometric objects, viz. nodes, edges, and polygon
package geo

import (
	"math"
	"polyGo/tools"
	"polyGo/vector"
)

// Polygon is a struct that collects nodes, edges, and centre of a polygon
type Polygon struct {
	id         int64
	Nodes      Nodes
	Edges      Edges
	Centre     vector.Vec
	isComplete bool
	Area       float64
}

// findCentre returns the geometric centre of a set of nodes
func findCentre(ns Nodes) vector.Vec {

	// sum of all vectors
	vSum := ns.vecs().Foldr(
		vector.Sum,
		vector.Zero,
	)

	// normalise
	s := 1 / float64(len(ns))
	return vector.ScalarMult(s, vSum)
}

// linkEdges returns edges for a set of nodes
func linkEdges(as Nodes, cent vector.Vec) Edges {

	// reordered as, first element last
	bs := append(as[1:], as[0])

	es := make(Edges, len(as))
	for i := range as {
		es[i] = NewEdge(as[i], bs[i], cent)
	}
	return es
}

func completeness(es Edges) bool {

	// the first and last node have a common node
	isClosed := es[0].tail == es[len(es)-1].head

	// all edges are oriented in same direction of rotation
	edgesOriented := tools.MapBoolFunToFloats(
		math.Signbit,
		es.dets(),
	)
	isOriented := tools.EachTrue(edgesOriented) || tools.EachFalse(edgesOriented)
	return isClosed && isOriented
}

// NewPoly forms a polygon struct from a unique set of Nodes
func NewPoly(nodes Nodes) Polygon {

	cent := findCentre(nodes)
	edges := linkEdges(nodes, cent)
	isComplete := completeness(edges)
	var area float64
	if isComplete {
		area = tools.FloatSum(edges.dets()) / 2
	} else {
		area = math.NaN()
	}

	return Polygon{
		id(),
		nodes,
		edges,
		cent,
		isComplete,
		area,
	}
}
