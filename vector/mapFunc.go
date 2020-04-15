// Package vector provides a type and functions for simple vector calculations
package vector

// mapFunc provides functions that

// MapScalarFunToVecs implements a specific case of list comprehension 'map'
// it takes a function (Vec -> float64) and applies it to each
// element of a slice of vectors. A slice of scalars is returned.
// MapFunToInts :: (Vec -> float64) -> [Vec] -> [float64]
func MapScalarFunToVecs(f func(Vec) float64, xs Vecs) []float64 {
	ys := make([]float64, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

// MapVecFunToVecs maps a function that takes and returns a vector
// to a slice of vectors.
// MapVecFunToVecs (Vec -> Vec) -> [Vec] -> [Vec]
func MapVecFunToVecs(f func(Vec) Vec, xs Vecs) Vecs {
	ys := make(Vecs, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

// FoldrVecs folds a function that takes and returns a vector
// to a slice of vectors.
// f accepts two arguments and returns one, all vectors
// xs are a slice of vectors
// x0 is the neutral element for the operation in f
// FoldrVecs (Vec -> Vec -> Vec) -> [Vec] -> Vec -> Vec
func foldrVecs(f func(Vec, Vec) Vec, xs Vecs, x0 Vec) Vec {
	y := x0
	for _, x := range xs {
		y = f(y, x)
	}
	return y
}

// mapVecFunToPoints maps a funct (Point -> Vec) to a slice of points
func mapVecFunToPoints(f func(p Point) Vec, xs Points) Vecs {
	ys := make(Vecs, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}
