package geobjects

import "polyGo/vector"

// helper functions to map functions to slices

// MapNodeFunToVecs maps a function that returns a Node to a slice of vector.Vec
func MapNodeFunToVecs(f func(vector.Vec) Node, xs []vector.Vec) []Node {
	ys := make([]Node, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}
