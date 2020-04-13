// Package tools provides useful functions for PolyGo
package tools

import (
	v "polyGo/vector"
)

// MapFunToInts implements a specific case of list comprehension 'map'
// it takes a function (uint32 -> uint32) and applies it to each
// element of a slice of uint. It's Hindley-Milner-Signature is:
// MapFunToInts :: (uint32 -> uint32) -> [uint32] -> [uint32]
func MapFunToInts(f func(uint32) uint32, xs []uint32) []uint32 {
	ys := make([]uint32, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

// MapScalarFunToVecs implements a specific case of list comprehension 'map'
// it takes a function (Vec -> float64) and applies it to each
// element of a slice of vectors. A slice of scalars is returned.
// It's Hindley-Milner-Signature is:
// MapFunToInts :: (Vec -> float64) -> [Vec] -> [float64]
func MapScalarFunToVecs(f func(v.Vec) float64, xs []v.Vec) []float64 {
	ys := make([]float64, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

// PointsToVecs convets a slice of PointXY to a slice of Vec
func PointsToVecs(xs []v.PointXY) []v.Vec {
	ys := make([]v.Vec, len(xs))
	for i, x := range xs {
		ys[i] = v.FromStruct(x)
	}
	return ys
}
