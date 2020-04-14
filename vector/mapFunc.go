// Package vector provides a type and functions for simple vector calculations
package vector

// mapFunc provides functions that

// MapScalarFunToVecs implements a specific case of list comprehension 'map'
// it takes a function (Vec -> float64) and applies it to each
// element of a slice of vectors. A slice of scalars is returned.
// MapFunToInts :: (Vec -> float64) -> [Vec] -> [float64]
func MapScalarFunToVecs(f func(Vec) float64, xs []Vec) []float64 {
	ys := make([]float64, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

// MapVecFunToVecs maps a function that takes and returns a vector
// to a slice of vectors.
// MapVecFunToVecs (Vec -> Vec) -> [Vec] -> [Vec]
func MapVecFunToVecs(f func(Vec) Vec, xs []Vec) []Vec {
	ys := make([]Vec, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

// PointsToVecs converts a slice of PointXY to a slice of Vec
func PointsToVecs(xs []PointXY) []Vec {
	ys := make([]Vec, len(xs))
	for i, x := range xs {
		ys[i] = FromStruct(x)
	}
	return ys
}
