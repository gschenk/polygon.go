package geo

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

// mapVecFunToNodes a function that returns a Vec to a slice of Nodes
func mapVecFunToNodes(f func(Node) vector.Vec, xs []Node) []vector.Vec {
	ys := make([]vector.Vec, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

// mapFloatFunToEdges a function that returns a Node to a slice of vector.Vec
func mapFloatFunToEdges(f func(Edge) float64, xs []Edge) []float64 {
	ys := make([]float64, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}
