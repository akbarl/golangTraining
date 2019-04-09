package point

type Point struct {
	X, Y int
}

func (p1 Point) DoDai(Point p2) float64 {
	return math.Sqrt((p1.X - p2.X) * (p))
}