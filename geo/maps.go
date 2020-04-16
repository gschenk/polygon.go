package geo

import "polyGo/vector"

// helper functions to map functions to slices

// mapVecFunToNodes a function that returns a Vec to a slice of Nodes
func mapVecFunToNodes(f func(Node) vector.Vec, xs []Node) vector.Vecs {
	ys := make(vector.Vecs, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

// mapIntFunToNodes a function that returns an int to a slice of Nodes
func mapIntFunToNodes(f func(Node) int64, xs []Node) []int64 {
	ys := make([]int64, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

// mapFloatFunToEdges maps a function that returns a float to each Edge
func mapFloatFunToEdges(f func(Edge) float64, xs []Edge) []float64 {
	ys := make([]float64, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

// mapBoolFunToEdges a function to Edges that returns a slice of bools
func mapBoolFunToEdges(f func(Edge) bool, xs []Edge) []bool {
	ys := make([]bool, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

// mapVecFunToEdges maps a function that returns a slice of Vecs to a slice of Edges
func mapVecsFunToEdges(f func(Edge) vector.Vecs, xs Edges) []vector.Vecs {
	ys := make([]vector.Vecs, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}
