package main

import "fmt"

type Equaler interface {
	Equal(Equaler) bool
}

type Point struct {
	X, Y int
}

func (p *Point) Equal(other Equaler) bool {

	q, ok := other.(*Point)
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

func (v *Vertex) Equal(other Equaler) bool {

	w, ok := other.(*Vertex)
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

func main() {

	var a, b, c, d Equaler

	a = &Point{12, 0}
	b = &Point{-2, 14}
	c = &Point{-2, 14}

	fmt.Println("a == b -", a.Equal(b))
	fmt.Println("b == c -", b.Equal(c))
	fmt.Println("c == a -", c.Equal(a))

	d = &Vertex{12, 0}

	fmt.Println("a == d -", a.Equal(d))
}
