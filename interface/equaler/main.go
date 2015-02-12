package main

import "fmt"

type Equaler interface {
	Equal(Equaler) bool
}

type Point struct {
	X, Y int
}

func (p *Point) Equal(e Equaler) bool {

	q, ok := e.(*Point)
	if !ok {
		return false
	}

	if p.X != q.X {
		return false
	}
	if p.Y != q.Y {
		return false
	}

	return true
}

type Vertex struct {
	X, Y int
}

func (v *Vertex) Equal(e Equaler) bool {

	w, ok := e.(*Vertex)
	if !ok {
		return false
	}

	if v.X != w.X {
		return false
	}
	if v.Y != w.Y {
		return false
	}

	return true
}

func equalSample(a, b Equaler) {

	var s string

	if a.Equal(b) {
		s = "=="
	} else {
		s = "!="
	}

	fmt.Printf("%v %s %v\n", a, s, b)
}

func main() {

	var a, b, c, d, e Equaler

	a = &Point{12, 0}
	b = &Point{-2, 14}
	c = &Point{-2, 14}
	d = &Vertex{12, 0}
	e = &Vertex{12, 0}

	equalSample(a, b)
	equalSample(b, c)
	equalSample(c, a)
	equalSample(a, d)
	equalSample(d, e)
}
