// Package extrema provides the function FindValues to find points in a point
// cloud in the Cartesian plane. The coordinates of these points have extreme
// values along the cardinal or diagonal axes.
//
package extrema

import (
	v "polyGo/vector"
	//    "polyGo/tools"
)

// accumulates vectors of extreme values at directions
// left, botom left, bottom, bottom right,
// right, top right, top, top left
type Accu [8]v.Vec

// compare function type
type compare func(v.Vec, v.Vec) bool

// A number of functions that return Bool are defined here.
// These functions compare two points given by their position vector.
// The naming follows the convention of a plot posted to a wall, where
// the x-axis increases to the right and the y-axis to the top.
//
// There is no ternary in Go!
// ordered array of comparison functions
var compFuncs = [8]compare{
	//	Left, BottomLeft, Bottom, BottomRight, Right, TopRight, Top, TopLeft
	func(a, b v.Vec) bool { return a[0] < b[0] },
	func(a, b v.Vec) bool { return a[0]+a[1] < b[0]+b[1] },
	func(a, b v.Vec) bool { return a[1] < b[1] },
	func(a, b v.Vec) bool { return a[0]-a[1] > b[0]-b[1] },
	func(a, b v.Vec) bool { return a[0] > b[0] },
	func(a, b v.Vec) bool { return a[0]+a[1] > b[0]+b[1] },
	func(a, b v.Vec) bool { return a[1] > b[1] },
	func(a, b v.Vec) bool { return a[0]-a[1] < b[0]-b[1] },
}

// tests two vectors with a test function and returns
// the first when true otherwise the second
func decide(a, b v.Vec, f compare) v.Vec {
	if f(a, b) {
		return a
	}
	return b
}

func makeInitAccu(a v.Vec) Accu {
	return [8]v.Vec{a, a, a, a, a, a, a, a}
}

// accu: stores the maximum or minium values found in tests with decide
// result: will be returned with changed max and uses as accu in next iteration
func innerFindValues(point v.Vec, accu Accu) Accu {
	var result Accu

	// loop over all 8 accu values using functions stored in compFuncs
	for i, old := range accu {
		result[i] = decide(point, old, compFuncs[i])
	}
	return result
}

// FindValues takes an initial point that is certainly within
// the point cloud and a slice of points. All points are position
// vectors of type Vec.
// The function returns up to eight unique position vectors for
// points with extreme coordinate values along cardinal and diagonal
// axes. The resulting slice may contain duplicates.
func FindValues(iniPoint v.Vec, points v.Vecs) v.Vecs {
	accu := makeInitAccu(iniPoint)
	for _, p := range points {
		accu = innerFindValues(p, accu)
	}
	extremes := make(v.Vecs, len(accu))
	for i, a := range accu {
		extremes[i] = a
	}
	return extremes
}
