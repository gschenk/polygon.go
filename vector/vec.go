package vector

// Vec type for a vector with two dimensions
type Vec [2]float64

// Vecs provides a slice of vectors Vec
type Vecs []Vec

// Foldr folds the slice with a function that returns a vector
// it needs an initial value v0 (eg neutral element of op)
// Foldr ::  (Vec -> Vec -> Vec) -> Vec -> Vec
func (vs Vecs) Foldr(f func(Vec, Vec) Vec, v0 Vec) Vec {
	return foldrVecs(f, vs, v0)
}
