// Package vector provides a type and functions for simple vector calculations
package vector

// Vec type for a vector with two dimensions
type Vec [2]float64

// Vecs provides a slice of vectors Vec
type Vecs []Vec

// Norm returns the norm of a vector
func (v Vec) Norm() float64 { return norm(v) }

// Norm2 returns the squared norm of a vector (dot product with itself)
func (v Vec) Norm2() float64 { return normSquare(v) }

// Foldr folds the slice with a function that returns a vector
// it needs an initial value v0 (eg neutral element of op)
// Foldr ::  (Vec -> Vec -> Vec) -> Vec -> Vec
func (vs Vecs) Foldr(f func(Vec, Vec) Vec, v0 Vec) Vec {
	return foldrVecs(f, vs, v0)
}

// Norms returns the norm of a vector
func (vs Vecs) Norms() []float64 {
	return mapScalarFunToVecs(
		func(v Vec) float64 { return v.Norm() },
		vs,
	)
}

// Norm2s returns the squared norm of a vector (dot product with itself)
func (vs Vecs) Norm2s() []float64 {
	return mapScalarFunToVecs(
		func(v Vec) float64 { return v.Norm2() },
		vs,
	)
}

// Undup returns Vecs without duplicates
func (vs Vecs) Undup() Vecs { return RemoveDuplicates(vs) }
