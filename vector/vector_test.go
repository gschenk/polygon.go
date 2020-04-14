package vector

import "testing"

// test vectors
var zero = Vec{0, 0}
var unitX = Vec{1, 0}
var unitY = Vec{0, 1}
var vecA = Vec{2, 1}
var vecB = Vec{1, 2}
var vecC = Vec{-3.4, 2.5}

func TestFromStruct(t *testing.T) {
	testPoint := PointXY{X: 1.01e3, Y: 0.02}
	result := FromStruct(testPoint)
	expected := Vec{1010, 0.02}

	if result != expected {
		t.Errorf(
			"PointXY incorrectly read as vector; expected %v, received %+v",
			expected,
			result,
		)
	}
}

func TestSum(t *testing.T) {
	tables := []struct {
		a Vec
		b Vec
		r Vec
	}{
		{vecA, vecA, Vec{4, 2}},
		{vecA, vecB, Vec{3, 3}},
		{vecB, vecA, Vec{3, 3}},
		{vecA, vecC, Vec{-1.4, 3.5}},
	}

	for _, table := range tables {
		result := Sum(table.a, table.b)
		if result != table.r {
			t.Errorf(
				"Vector sum calculated wrongly; expected %v, received %v",
				table.r,
				result,
			)
		}
	}
}

func TestFromAtoB(t *testing.T) {
	tables := []struct {
		a Vec
		b Vec
		r Vec
	}{
		{zero, vecA, vecA},
		{vecA, vecA, zero},
		{vecA, vecB, Vec{-1, 1}},
		{vecB, vecA, Vec{1, -1}},
		{vecA, vecC, Vec{-5.4, 1.5}},
	}

	for _, table := range tables {
		result := FromAtoB(table.a, table.b)
		if result != table.r {
			t.Errorf(
				"Vector TestFromAtoB calculated wrongly; expected %v, received %v",
				table.r,
				result,
			)
		}
	}
}

func TestDot(t *testing.T) {
	tables := []struct {
		a Vec
		b Vec
		x float64
	}{
		{unitX, unitY, 0}, // orthogonal vectors zero
		{vecA, vecA, 5},
		{vecA, vecB, 4},
		{vecB, vecA, 4},
		{vecA, vecC, -4.3},
	}

	for _, table := range tables {
		result := Dot(table.a, table.b)
		if result != table.x {
			t.Errorf(
				"Dot product calculated wrongly; expected %f, received %f",
				table.x,
				result,
			)
		}
	}
}

func TestDet(t *testing.T) {
	tables := []struct {
		a Vec
		b Vec
		x float64
	}{
		{unitX, unitY, 1},     //orthogonal vectors: product of norms
		{vecA, vecA, 0},       // identical vectors: zero
		{unitX, Vec{2, 0}, 0}, // collinear vectors: zero
		{vecA, vecB, 3},
		{vecB, vecA, -3},
		{vecA, vecC, 8.4},
	}

	for _, table := range tables {
		result := Det(table.a, table.b)
		if result != table.x {
			t.Errorf(
				"Determinant calculated wrongly; expected %f, received %f",
				table.x,
				result,
			)
		}
	}
}

func TestNormSquare(t *testing.T) {
	tables := []struct {
		a Vec
		x float64
	}{
		{zero, 0},
		{unitX, 1},
		{unitY, 1},
		{vecA, 5},
		{vecB, 5},
		{vecC, 17.81},
	}

	for _, table := range tables {
		result := NormSquare(table.a)
		if result != table.x {
			t.Errorf(
				"Norm square calculated wrongly; expected %f, received %f",
				table.x,
				result,
			)
		}
	}
}
