// Package tools provides useful functions for PolyGo
package tools

import "math"

// Epsilon provides machine epsilon for float64 around 1
var Epsilon = math.Nextafter(1, 2) - 1

// FloatIsZero returns true if a real value is within an epsilon neighbourhood of zero
func FloatIsZero(x float64) bool {
	return math.Abs(x) < 2*Epsilon
}

// FloatsEqual returns true if the difference of two floats is within an
// epsilon neighbourhood of zero
func FloatsEqual(x, y float64) bool {
	return FloatIsZero(x - y)
}
