// Package vector provides a type and functions for simple vector calculations
package vector

import (
	"math"
	"polyGo/tools"
)

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

// SafeFromAtoB is FromAtoB with checks for equal component values
// returns second argument, false when both vectors are same
// SafeFromAtoB :: Vec -> Vec -> Vec -> bool
func SafeFromAtoB(a, b Vec) Vec {
	xIsZero := tools.FloatsEqual(a[0], b[0])
	yIsZero := tools.FloatsEqual(a[1], b[1])
	if xIsZero && yIsZero {
		return Zero
	}
	if xIsZero {
		y := b[1] - a[1]
		return Vec{0, y}
	}
	if yIsZero {
		x := b[0] - a[0]
		return Vec{x, 0}
	}
	x := b[0] - a[0]
	y := b[1] - a[1]
	return Vec{x, y}
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
// Det :: Vec -> Vec -> float64
func Det(a, b Vec) float64 {
	return a[0]*b[1] - a[1]*b[0]
}

// SafeDet is Det with a check when close to zero
// Det :: Vec -> Vec -> float64
func SafeDet(a, b Vec) float64 {
	down := a[0] * b[1]
	up := a[1] * b[0]
	if tools.FloatsEqual(down, up) {
		return 0
	}
	return down - up
}

// normSquare returns the squared norm of a vector (dot product with itself)
// normSquare :: Vec -> float64
func normSquare(a Vec) float64 {
	return math.Pow(a[0], 2) + math.Pow(a[1], 2)
}

// norm returns the norm of a vector
// norm :: Vec -> float64
func norm(a Vec) float64 {
	return math.Sqrt(normSquare(a))
}

// EqualVecs compares two vectors, ie difference of components less than epsilon
// EqualVecs :: Vec -> Vec -> bool
func EqualVecs(a, b Vec) bool {
	return tools.FloatsEqual(a[0], b[0]) &&
		tools.FloatsEqual(a[1], b[1])
}
