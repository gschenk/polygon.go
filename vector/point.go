package vector

// Point type for a point in a Carthesian plane
type Point struct {
	X float64
	Y float64
}

// Points provides a slice of Point
type Points []Point

// vec returns the position vector
func (p Point) vec() Vec {
	return Vec{p.X, p.Y}
}

// Vecs returns a slice of the Points' position vectors
func (ps Points) Vecs() Vecs {
	return mapVecFunToPoints(
		func(p Point) Vec { return p.vec() },
		ps,
	)
}
