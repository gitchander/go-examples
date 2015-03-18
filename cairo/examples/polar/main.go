package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"

	"github.com/gitchander/go-examples/cairo"
)

type Decart struct {
	X, Y float64
}

func (a Decart) Scale(s float64) Decart {
	return Decart{a.X * s, a.Y * s}
}

func (a Decart) Add(b Decart) Decart {
	return Decart{a.X + b.X, a.Y + b.Y}
}

func (a Decart) Sub(b Decart) Decart {
	return Decart{a.X - b.X, a.Y - b.Y}
}

type Polar struct {
	R, Phi float64
}

func PolarToDecart(polar Polar) Decart {

	sinPhi, cosPhi := math.Sincos(polar.Phi)

	x := polar.R * cosPhi
	y := polar.R * sinPhi

	return Decart{x, y}
}

type Range struct {
	Min   float64
	Max   float64
	Count int
}

func (r *Range) Step() float64 {
	return (r.Max - r.Min) / float64(r.Count-1)
}

type PolarCurve struct {
	Name          string
	Scale         float64
	RadiusByAngle func(Angle float64) float64
	Ranges        []Range
}

var ps = []PolarCurve{
	PolarCurve{
		Name:  "spiral",
		Scale: 7,
		RadiusByAngle: func(angle float64) float64 {
			return angle
		},
		Ranges: []Range{
			Range{Min: 0, Max: math.Pi * 10, Count: 500},
		},
	},
	PolarCurve{
		Name:  "cardioid",
		Scale: 200,
		RadiusByAngle: func(angle float64) float64 {
			return math.Sin(angle / 2)
		},
		Ranges: []Range{
			Range{Min: 0, Max: math.Pi * 2, Count: 100},
		},
	},
	PolarCurve{
		Name:  "lemniscate",
		Scale: 2,
		RadiusByAngle: func(angle float64) float64 {
			const a = 100
			c := math.Cos(2 * angle)
			if c < 0 {
				return 0
			}
			return math.Sqrt(a * a * c)
		},
		Ranges: []Range{
			Range{Min: -math.Pi / 4, Max: math.Pi / 4, Count: 50},
			Range{Min: math.Pi * 3 / 4, Max: math.Pi * 5 / 4, Count: 50},
		},
	},
	PolarCurve{
		Name:  "cannabis",
		Scale: 200,
		RadiusByAngle: func(angle float64) float64 {
			a := 0.3
			return a * (1.0 + 9.0/10.0*math.Cos(8.0*angle)) * (1.0 + 1.0/10.0*math.Cos(24.0*angle)) * (9.0/10.0 + 1.0/10.0*math.Cos(200.0*angle)) * (1.0 + math.Sin(angle))
		},
		Ranges: []Range{
			Range{Min: 0, Max: math.Pi * 2, Count: 1000},
		},
	},
	PolarCurve{
		Name:  "ellipse",
		Scale: 200,
		RadiusByAngle: func(angle float64) float64 {

			var (
				a = 0.7
				e = 0.6
				l = a * (1 - e*e)
			)

			return l / (1.0 - e*math.Cos(angle))
		},
		Ranges: []Range{
			Range{Min: 0, Max: math.Pi * 2, Count: 100},
		},
	},
	PolarCurve{
		Name:  "custom",
		Scale: 200,
		RadiusByAngle: func(angle float64) float64 {
			return math.Sin(angle * 3)
		},
		Ranges: []Range{
			Range{Min: 0, Max: math.Pi, Count: 100},
		},
	},
}

func makeDir(dir string) error {

	fi, err := os.Stat(dir)
	if err != nil {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return err
		}
	} else {
		if !fi.IsDir() {
			return errors.New("file is not dir")
		}
	}
	return nil
}

func Create(dir string, params PolarCurve) {

	width, height := 512, 512

	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, width, height)
	canvas, _ := cairo.NewCanvas(surface)

	canvas.SetLineJoin(cairo.LINE_JOIN_ROUND)

	canvas.SetSourceRGB(1.0, 1.0, 1.0)
	canvas.Rectangle(0.0, 0.0, float64(width), float64(height))
	canvas.Fill()

	canvas.SetLineWidth(1.0)
	canvas.SetSourceRGB(0.7, 0.7, 0.7)
	DrawAxes(canvas, width, height)

	canvas.SetLineWidth(2)
	canvas.SetSourceRGB(0.5, 0, 0)
	PolarDraw(canvas, width, height, params)

	fileName := filepath.Join(dir, fmt.Sprintf("%s.png", params.Name))

	surface.WriteToPNG(fileName)
	surface.Finish()
}

func DrawAxes(canvas *cairo.Canvas, width, height int) {

	var (
		x0 = float64(width) * 0.5
		y0 = float64(height) * 0.5
	)

	const rd = float64(40)
	k := 6
	m := 80
	du := 2 * math.Pi / float64(m-1)

	for i := 0; i < k; i++ {

		u := float64(0)
		for j := 0; j < m; j++ {

			s, c := math.Sincos(u)

			r := rd * float64(i+1)
			x := x0 + r*c
			y := y0 + r*s

			if j == 0 {
				canvas.MoveTo(x, y)
			} else {
				canvas.LineTo(x, y)
			}

			u += du
		}
	}

	n := 16
	du = 2 * math.Pi / float64(n)
	u := float64(0)

	for i := 0; i < n; i++ {

		s, c := math.Sincos(u)

		r := rd * float64(1)
		x1 := x0 + r*c
		y1 := y0 + r*s

		r = rd * float64(k)
		x2 := x0 + r*c
		y2 := y0 + r*s

		canvas.MoveTo(x1, y1)
		canvas.LineTo(x2, y2)

		u += du
	}

	// render center cross
	{
		canvas.MoveTo(x0, y0-rd)
		canvas.LineTo(x0, y0+rd)

		canvas.MoveTo(x0-rd, y0)
		canvas.LineTo(x0+rd, y0)
	}

	canvas.Stroke()
}

func PolarDraw(canvas *cairo.Canvas, width, height int, params PolarCurve) {

	var center = Decart{
		X: float64(width) * 0.5,
		Y: float64(height) * 0.5,
	}

	for _, r := range params.Ranges {

		angleStep := r.Step()
		n := r.Count
		p := Polar{Phi: r.Min}

		for i := 0; i < n; i++ {

			p.R = params.RadiusByAngle(p.Phi)
			d := PolarToDecart(p)
			p.Phi += angleStep

			temp := d.Scale(params.Scale)
			temp = center.Add(temp)

			x, y := temp.X, temp.Y

			if i == 0 {
				canvas.MoveTo(x, y)
			} else {
				canvas.LineTo(x, y)
			}
		}

		canvas.Stroke()
	}
}

func main() {

	path := os.Args[0]
	dir, _ := filepath.Split(path)
	dir = filepath.Join(dir, "result")

	err := makeDir(dir)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, p := range ps {
		Create(dir, p)
	}
}
