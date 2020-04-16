// Package tools provides useful functions for PolyGo
package tools

import "math"

const tol = 8 // 2 exponent multiplied with epsilon, tune to desired precision

// Epsilon provides machine epsilon for float64 around 1
var Epsilon = math.Exp2(tol) * (math.Nextafter(1, 2) - 1)
var Theta = math.Exp2(tol) * math.SmallestNonzeroFloat64

// FloatIsZero returns true if a real value is within an epsilon neighbourhood of zero
func FloatIsZero(x float64) bool {
	return math.Abs(x) < Epsilon
}

// FloatsEqual returns true if the difference of two floats is within an
// epsilon neighbourhood of zero
// Source https://stackoverflow.com/a/32334103/3842889
func FloatsEqual(x, y float64) bool {
	if x == y {
		return true
	}
	diff := math.Abs(x - y)
	norm := math.Abs(x) + math.Abs(y)
	return diff < math.Max(norm*Epsilon, Theta)
}
