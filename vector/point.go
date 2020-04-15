package vector

// Point type for a point in a Carthesian plane
type Point struct {
	X float64
	Y float64
}

// vec returns the position vector
func (p Point) vec() Vec {
	return Vec{p.X, p.Y}
}
