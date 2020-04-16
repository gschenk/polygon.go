package geo

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
