package main

import "fmt"

type Optional[T any] struct {
	Present bool
	Value   T
}

// NewPresent
// MakePresent
func PresentValue[T any](val T) Optional[T] {
	return Optional[T]{
		Present: true,
		Value:   val,
	}
}

func (v Optional[T]) GetValue() (T, bool) {
	if v.Present {
		return v.Value, true
	}
	var val T
	return val, false
}

func (p *Optional[T]) SetValue(val T) {
	*p = PresentValue(val)
}

func (p *Optional[T]) Reset() {
	*p = Optional[T]{}
}

func (v Optional[T]) If(f func(T)) {
	if v.Present {
		f(v.Value)
	}
}

func testOptional() {
	var o Optional[float64]
	o = PresentValue[float64](23)
	o.SetValue(-7)
	//o.Reset()
	fmt.Println(o)
}
