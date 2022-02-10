package devilc

type Point2f struct {
	X, Y float64
}

func (p Point2f) Add(q Point2f) Point2f {
	return Point2f{
		X: p.X + q.X,
		Y: p.Y + q.Y,
	}
}

func (p Point2f) Sub(q Point2f) Point2f {
	return Point2f{
		X: p.X - q.X,
		Y: p.Y - q.Y,
	}
}

func (p Point2f) MulScalar(k float64) Point2f {
	return Point2f{
		X: p.X * k,
		Y: p.Y * k,
	}
}

func (p Point2f) DivScalar(k float64) Point2f {
	return Point2f{
		X: p.X / k,
		Y: p.Y / k,
	}
}
