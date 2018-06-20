package life

import (
	"github.com/gitchander/go-examples/gif/gol/bitmap"
)

type Torus struct {
	size       Point
	curr, next *bitmap.Bitmap
}

func NewTorus(size Point) *Torus {
	if size.X < 0 {
		size.X = 0
	}
	if size.Y < 0 {
		size.Y = 0
	}
	n := size.X * size.Y
	return &Torus{
		size: size,
		curr: bitmap.New(n),
		next: bitmap.New(n),
	}
}

func (t *Torus) Size() (x, y int) {
	return t.size.X, t.size.Y
}

func (t *Torus) Empty() bool {
	return (t.size.X == 0) || (t.size.Y == 0)
}

func (t *Torus) swapBuffers() {
	t.curr, t.next = t.next, t.curr
}

func (t *Torus) Next() {
	if t.Empty() {
		return
	}
	for y := 0; y < t.size.Y; y++ {
		for x := 0; x < t.size.X; x++ {
			var (
				i = t.offset(x, y)
				n = t.neighbors(x, y)
			)
			v, _ := t.curr.Get(i)
			if v {
				v = (n == 2) || (n == 3)
			} else {
				v = (n == 3)
			}
			t.next.Set(i, v)
		}
	}
	t.swapBuffers()
}

func (t *Torus) neighbors(x, y int) (n int) {
	for _, ns := range neShifts {
		offset := t.offset(x+ns.X, y+ns.Y)
		v, _ := t.curr.Get(offset)
		if v {
			n++
		}
	}
	return
}

func (t *Torus) offset(x, y int) int {
	x = mod(x, t.size.X)
	y = mod(y, t.size.Y)
	return y*t.size.X + x
}

func mod(x, y int) int {
	t := x % y
	if t < 0 {
		t += y
	}
	return t
}

func (t *Torus) Get(x, y int) bool {
	if t.Empty() {
		return false
	}
	v, _ := t.curr.Get(t.offset(x, y))
	return v
}

func (t *Torus) Set(x, y int, v bool) {
	if t.Empty() {
		return
	}
	t.curr.Set(t.offset(x, y), v)
}

func (t *Torus) Clone() interface{} {
	c := NewTorus(t.size)
	n := t.size.X * t.size.Y
	for i := 0; i < n; i++ {
		v, _ := t.curr.Get(i)
		c.curr.Set(i, v)
	}
	return c
}

func (t1 *Torus) Equal(a interface{}) bool {

	t2, ok := a.(*Torus)
	if !ok {
		return false
	}

	if (t1.size.X != t2.size.X) || (t1.size.Y != t2.size.Y) {
		return false
	}
	n := t1.size.X * t1.size.Y
	for i := 0; i < n; i++ {
		v1, _ := t1.curr.Get(i)
		v2, _ := t2.curr.Get(i)
		if v1 != v2 {
			return false
		}
	}
	return true
}
