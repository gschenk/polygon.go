package vector

import "testing"

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

func TestDet(t *testing.T) {
	tables := []struct {
		a Vec
		b Vec
		x float64
	}{
		{
			Vec{2, 1},
			Vec{2, 1},
			0,
		},
		{
			Vec{2, 1},
			Vec{1, 2},
			3,
		},
		{
			Vec{1, 2},
			Vec{2, 1},
			-3,
		},
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
