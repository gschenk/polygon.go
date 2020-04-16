// Package geo provides geometric objects, viz. nodes, edges, and polygon
package geo

import "polyGo/vector"

// Node is a struct that provides the position vector of a node of a polygon
type Node struct {
	id int64
	v  vector.Vec
}

// Nodes provides a slice of Node
type Nodes []Node

// NewNode creates a node struct from a position vector
func NewNode(a vector.Vec) Node {
	return Node{id(), a}
}

// NewNodes constructs a slice of Node from a slice of Vec
func NewNodes(vs vector.Vecs) Nodes {
	ns := make([]Node, len(vs))
	for i, v := range vs {
		ns[i] = NewNode(v)
	}
	return ns
}

// vecs returns a slice of Node's position vectors
func (ns Nodes) vecs() vector.Vecs {
	return mapVecFunToNodes(
		func(n Node) vector.Vec { return n.v },
		ns,
	)
}

// vecs returns a slice of Node's ids
func (ns Nodes) Ids() []int64 {
	return mapIntFunToNodes(
		func(n Node) int64 { return n.id },
		ns,
	)
}
