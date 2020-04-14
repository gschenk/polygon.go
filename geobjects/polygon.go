// Package geobjects provides geometric objects, viz. nodes, edges, and polygon
package geobjects

import "polyGo/vector"

// Polygon is a struct that collects nodes, edges, and centre of a polygon
type Polygon struct {
	centre   vector.Vec
	complete bool
}

// MakePoly forms a polygon struct from a unique set of points given as position vectors
func MakePoly([]vector.Vec) Polygon {
	return Polygon{complete: false} //stub
}
