// Package tools provides useful functions for PolyGo
package tools

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
