// Package vector provides a type and functions for simple vector calculations
package vector

import (
	"reflect"
	"testing"
)

func TestMapScalarFunToVecs(t *testing.T) {

	f := func(a Vec) float64 { return a[0] }
	expected := []float64{0, 1, 0, 2, 1, -3.4}
	result := MapScalarFunToVecs(f, testVs)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			"Test func not mapped correctly to testVs; expected %v, received %v",
			expected,
			result,
		)
	}
}

func TestMapVecFunToVecs(t *testing.T) {

	f := func(a Vec) Vec { return a }
	expected := testVs
	result := MapVecFunToVecs(f, testVs)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			"Test func not mapped correctly to testVs; expected %v, received %v",
			expected,
			result,
		)
	}
}

func TestFoldrVecs(t *testing.T) {

	f := Sum
	xs := testVs
	x0 := Zero
	expected := Vec{0.6, 6.5}
	result := FoldrVecs(f, xs, x0)

	if !EqualVecs(expected, result) {
		t.Errorf(
			"Test func not folded correctly with Sum; expected %v, received %v",
			expected,
			result,
		)
	}
}

//(f func(Vec, Vec) Vec, xs []Vec, x0 Vec) Vec {
//	y := x0
//	for _, x := range xs {
//		y = f(y, x)
//	}
//	return y
//}
