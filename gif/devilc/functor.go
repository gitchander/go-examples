package devilc

// type CalcFunc func(t float64) Point2f

type Functor interface {
	Func(t float64) (Point2f, bool)
}
