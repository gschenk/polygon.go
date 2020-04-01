package vector


type PointXY struct{
    X float64
    Y float64
}

type Vec [2]float64


func FromStruct(r PointXY)  Vec {
    return Vec{r.X, r.Y}
}


// Det returns the determinant of a matrix with the elemts of two column vectors.
func Det(a, b Vec) float64 {
	return a[0]*b[1] - a[1]*b[0]
}

