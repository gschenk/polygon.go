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

// there is no ternary!
var fLeft = func (a, b v.Vec) bool {return a[0] < b[0]}
var fBotLeft = func (a, b v.Vec) bool {return a[0] + a[1] < b[0] + b[1]}
var fBot = func (a, b v.Vec) bool {return a[1] < b[1]}
var fBotRight = func (a, b v.Vec) bool {return a[0] - a[1] > b[0] - b[1]}
var fRight = func (a, b v.Vec) bool {return a[0] > b[0]}
var fTopRight = func (a, b v.Vec) bool {return a[0] + a[1] > b[0] + b[1]}
var fTop = func (a, b v.Vec) bool {return a[1] > b[1]}
var fTopLeft = func (a, b v.Vec) bool {return a[0] - a[1] < b[0] - b[1]}

// ordered array of these functions
var fs = [8]compare{fLeft, fBotLeft, fBot, fBotRight, fRight, fTopRight, fTop, fTopLeft}
// i cannot be bothered to write an array of functions to iterate over
// may i rot in hell for copy and paste


// tests two vectors with a test function and returns
// the first when true otherwise the second
func decide (a v.Vec, b v.Vec, f compare) v.Vec {
    var c v.Vec
    if f(a, b) {
        c = a
    } else {
        c = b
    }
    return c
}

func makeInitAccu (a v.Vec) Accu {
    return [8]v.Vec{a,a,a,a,a,a,a,a}
}

// accu: stores the maximum or minium values found in tests with decide
// result: will be returned with changed max and uses as accu in next iteration
func innerFindValues(point v.Vec, accu Accu) Accu {
    var result Accu

    // loop over all 8 accu values using functions stored in fs
    for i, old := range accu {
        result[i] = decide(point, old, fs[i])
    }
    return result
}

func FindValues (iniPoint v.Vec, points []v.Vec) []v.Vec{
    accu := makeInitAccu(iniPoint)
    for _, p := range points {
        accu = innerFindValues(p, accu)
    }
    extremes := make([]v.Vec, len(accu))
    for i, a := range accu {
        extremes[i] = a
    }
    return extremes
}

