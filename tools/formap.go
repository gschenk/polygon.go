package tools

import (
	v "polyGo/vector"
)

func MapFunToInts(f func(uint32) uint32, xs []uint32) []uint32 {
	ys := make([]uint32, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

func MapScalarFunToVecs(f func(v.Vec) float64, xs []v.Vec) []float64 {
	ys := make([]float64, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

func PointsToVecs(xs []v.PointXY) []v.Vec {
	ys := make([]v.Vec, len(xs))
	for i, x := range xs {
		ys[i] = v.FromStruct(x)
	}
	return ys
}
