package extrema

import (
	"polyGo/tools"
	v "polyGo/vector"
)

//// secondary strategy
// if only two extreme points are found this looks for up to two
// more points in very narrow shapes. If all collinear, it continues with
// two points.

// SecondExtremes may find up to two more extreme points it takes two previous
// extreme points, a, b, and a slice of points
func SecondExtremes(points v.Vecs, a, b v.Vec) v.Vecs {
	ab := v.FromAtoB(a, b)

	min := a
	minDet := 0.0
	max := a
	maxDet := 0.0

	for _, p := range points {
		ap := v.FromAtoB(a, p)
		det := v.SafeDet(ab, ap)

		// ignore colinear points
		if tools.FloatIsZero(det) {
			continue
		}

		if det < minDet {
			minDet = det
			min = p
		}
		if det > maxDet {
			maxDet = det
			max = p
		}
	}
	return v.Vecs{a, min, b, max}.Undup()
}
