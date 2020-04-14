// Package geobjects provides geometric objects, viz. nodes, edges, and polygon
package geobjects

import "polyGo/vector"

// Polygon is a struct that collects nodes, edges, and centre of a polygon
type Polygon struct {
	id       int64
	nodes    []Node
	centre   vector.Vec
	complete bool
}

// NewPoly forms a polygon struct from a unique set of Nodes
func NewPoly(ns []Node) Polygon {
	return Polygon{id: id(), nodes: ns, complete: false} //stub
}
