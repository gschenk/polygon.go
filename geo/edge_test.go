package geo

import (
	"polyGo/vector"
	"testing"
)

func TestProcessPoint(t *testing.T) {
	cent := vector.Zero
	nodeA := NewNode(vector.Vec{1, 1})
	nodeB := NewNode(vector.Vec{-1, 1})
	pointC := vector.Vec{0, 0.5}   // inside
	pointD := vector.Vec{0, 1}     // on edge
	pointE := vector.Vec{0, 1.5}   // Outside
	pointF := vector.Vec{0.5, 1.5} // Outside

	tables := []struct {
		a             Node
		b             Node
		p             vector.Vec
		q             vector.Vec
		expReturnFlag bool
		expNextExists bool
		expNextAngle  float64
		expLen        int
	}{
		{nodeA, nodeB, pointD, pointC, true, false, 0, 0},
		{nodeA, nodeB, pointC, pointD, false, false, 0, 0},
		{nodeA, nodeB, pointE, pointD, false, true, 0.8, 1},
		{nodeA, nodeB, pointD, pointD, false, false, 0, 0},
		{nodeA, nodeB, pointE, pointF, false, true, 2, 2},
		{nodeB, nodeA, pointD, pointC, true, false, 0, 0}, //reverse nodes
		{nodeB, nodeA, pointC, pointD, false, false, 0, 0},
		{nodeB, nodeA, pointD, pointD, false, false, 0, 0},
		{nodeB, nodeA, pointE, pointF, false, true, 0.8, 2}, //other Outside point found first
	}

	for _, table := range tables {

		// construct test edge
		edge := NewEdge(table.a, table.b, cent)

		// test point against edge
		edge.processPoint(table.p)

		_, rFlag := edge.processPoint(table.q)

		if rFlag != table.expReturnFlag {
			t.Errorf(
				"Inside/Outside point %v not recognised",
				table.p,
			)
		}
		if edge.HasOutside != table.expNextExists {
			t.Errorf(
				"Next node %v not recognised",
				table.p,
			)
		}
		if edge.nextAngle != table.expNextAngle {
			t.Errorf(
				"Next angle %f not correct, expected %f",
				edge.nextAngle,
				table.expNextAngle,
			)
		}
		if len(edge.Outside) != table.expLen {
			t.Errorf(
				"Length of Outside points slice not correct, expected %d, found %d",
				table.expLen,
				len(edge.Outside),
			)
		}
	}
}
