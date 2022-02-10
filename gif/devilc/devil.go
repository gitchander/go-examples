package devilc

import "math"

// Devil's Curve

// https://en.wikipedia.org/wiki/Devil%27s_curve
// https://mathworld.wolfram.com/DevilsCurve.html

type Devil struct {
	A float64
	B float64
}

func (d Devil) Functor() Functor {
	return &devilFunctor{
		aa: square(d.A),
		bb: square(d.B),
	}
}

type devilFunctor struct {
	aa float64
	bb float64
}

func (p *devilFunctor) Func(t float64) (Point2f, bool) {

	sin, cos := math.Sincos(t)

	var (
		s2 = sin * sin
		c2 = cos * cos
	)

	d := ((p.aa * s2) - (p.bb * c2)) / (s2 - c2)
	if d < 0 {
		return Point2f{}, false
	}

	sqrt := math.Sqrt(d)

	r := Point2f{
		X: cos * sqrt,
		Y: sin * sqrt,
	}

	return r, true
}
