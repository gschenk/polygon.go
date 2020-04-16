// Package vector provides a type and functions for simple vector calculations
package vector

// mapFunc provides functions that iterate through slices and apply functions
// element wise

// mapScalarFunToVecs implements a specific case of list comprehension 'map'
// it takes a function (Vec -> float64) and applies it to each
// element of a slice of vectors. A slice of scalars is returned.
// mapFunToInts :: (Vec -> float64) -> [Vec] -> [float64]
func mapScalarFunToVecs(f func(Vec) float64, xs Vecs) []float64 {
	ys := make([]float64, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

// foldrVecs folds a function that takes and returns a vector
// to a slice of vectors.
// f accepts two arguments and returns one, all vectors
// xs are a slice of vectors
// x0 is the neutral element for the operation in f
// foldrVecs (Vec -> Vec -> Vec) -> [Vec] -> Vec -> Vec
func foldrVecs(f func(Vec, Vec) Vec, xs Vecs, x0 Vec) Vec {
	y := x0
	for _, x := range xs {
		y = f(y, x)
	}
	return y
}
