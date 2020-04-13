// Package vector provides a type and functions for simple vector calculations
package vector

// PointXY type for a point in a Carthesian plane
type PointXY struct {
	X float64
	Y float64
}

// Vec type for a vector with two dimensions
type Vec [2]float64

// FromStruct returns the position vector of a coordinate in the Cartesian
// plane. It takes type PointXY and returns Vec
func FromStruct(r PointXY) Vec {
	return Vec{r.X, r.Y}
}

// Det returns the determinant of a matrix formed by the elements
// of two column vectors provied as Vec.
func Det(a, b Vec) float64 {
	return a[0]*b[1] - a[1]*b[0]
}
