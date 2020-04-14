// Package vector provides a type and functions for simple vector calculations
package vector

import (
	"math"
	"polyGo/tools"
)

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

// Sum returns the sum of two vectors
func Sum(a, b Vec) Vec {
	return Vec{
		a[0] + b[0],
		a[1] + b[1],
	}
}

// FromAtoB returns a vector from point A to point B where these points are
// given by their position vectors
// FromAtoB :: Vec -> Vec -> Vec
func FromAtoB(a, b Vec) Vec {
	return Vec{
		b[0] - a[0],
		b[1] - a[1],
	}
}

// ScalarMult returns the product of a scalar with a vector
// ScalarMult :: float64 -> Vec -> Vec
func ScalarMult(x float64, a Vec) Vec {
	return Vec{
		x * a[0],
		x * a[1],
	}
}

// Dot returns the dot product of two vectors, a and b
// Dot :: Vec -> Vec -> float64
func Dot(a, b Vec) float64 {
	return a[0]*b[0] + a[1]*b[1]
}

// Det returns the determinant of a matrix formed by the elements
// of two column vectors provied as Vec.
// Dot :: Vec -> Vec -> float64
func Det(a, b Vec) float64 {
	return a[0]*b[1] - a[1]*b[0]
}

// NormSquare returns the squared norm of a vector (dot product with itself)
// NormSquare :: Vec -> float64
func NormSquare(a Vec) float64 {
	return math.Pow(a[0], 2) + math.Pow(a[1], 2)
}

// Norm returns the norm of a vector
// Norm :: Vec -> float64
func Norm(a Vec) float64 {
	return math.Sqrt(NormSquare(a))
}

// EqualVecs compares two vectors, ie difference of components less than epsilon
// EqualVecs :: Vec -> Vec -> bool
func EqualVecs(a, b Vec) bool {
	return tools.FloatsEqual(a[0], b[0]) &&
		tools.FloatsEqual(a[1], b[1])
}
