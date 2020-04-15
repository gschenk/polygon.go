// Package tools provides useful functions for PolyGo
package tools

// MapBoolFunToFloats implements a specific case of list comprehension 'map'
// it takes a function (float64 -> bool) and applies it to each
// element of a slice. It's Hindley-Milner-Signature is:
// MapBoolFunToFloats:: (float64 -> bool) -> [float64] -> [bool]
func MapBoolFunToFloats(f func(float64) bool, xs []float64) []bool {
	ys := make([]bool, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

// FloatSum sums the values of a slice of float
// FloatSum :: [float64] -> float64
func FloatSum(xs []float64) float64 {
	var y float64 = 0 // neutral element
	for _, x := range xs {
		y += x
	}
	return y
}

// foldrBool :: (bool -> bool) -> [bool] -> bool
func foldrBool(f func(bool, bool) bool, xs []bool, x0 bool) bool {
	y := x0 // neutral element
	for _, x := range xs {
		y = f(y, x)
	}
	return y
}

// EachTrue returns true when each element of a slice is true
// EachTrue :: [bool] -> bool
func EachTrue(bs []bool) bool {
	return foldrBool(
		func(a, b bool) bool { return a && b },
		bs,
		true,
	)
}

// EachFalse returns false when each element of a slice is false
// EachFalse :: [bool] -> bool
func EachFalse(bs []bool) bool {
	return foldrBool(
		func(a, b bool) bool { return a && !b },
		bs,
		true,
	)
}
