package vector

import "testing"

func TestPoint(t *testing.T) {
	testPoint := Point{X: 1.01e3, Y: 0.02}
	result := testPoint.vec()
	expected := Vec{1010, 0.02}

	if result != expected {
		t.Errorf(
			"Point incorrectly read as vector; expected %v, received %+v",
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
		{testA, testA, Vec{4, 2}},
		{testA, testB, Vec{3, 3}},
		{testB, testA, Vec{3, 3}},
		{testA, testC, Vec{-1.4, 3.5}},
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
		{Zero, testA, testA},
		{testA, testA, Zero},
		{testA, testB, Vec{-1, 1}},
		{testB, testA, Vec{1, -1}},
		{testA, testC, Vec{-5.4, 1.5}},
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
		{UnitX, UnitY, 0}, // orthogonal vectors Zero
		{testA, testA, 5},
		{testA, testB, 4},
		{testB, testA, 4},
		{testA, testC, -4.3},
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
		{UnitX, UnitY, 1},     //orthogonal vectors: product of norms
		{testA, testA, 0},     // identical vectors: zero
		{UnitX, Vec{2, 0}, 0}, // collinear vectors: zero
		{testA, testB, 3},
		{testB, testA, -3},
		{testA, testC, 8.4},
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
		{Zero, 0},
		{UnitX, 1},
		{UnitY, 1},
		{testA, 5},
		{testB, 5},
		{testC, 17.81},
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
