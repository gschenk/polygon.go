// Package geobjects provides geometric objects, viz. nodes, edges, and polygon
package geobjects

import "polyGo/vector"

// Node is a struct that provides the position vector of a node of a polygon
type Node struct {
	id int64
	v  vector.Vec
}

// NewNode creates a node struct from a position vector
func NewNode(a vector.Vec) Node {
	return Node{id(), a}
}
