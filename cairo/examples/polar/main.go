package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"

	"github.com/gitchander/go-examples/cairo"
)

var ps = []PolarParams{
	PolarParams{
		Name:  "spiral",
		Scale: 5,
		GetRadius: func(angle float64) float64 {
			return angle
		},
		Angle: AngleSets{
			Min:  0,
			Max:  math.Pi * 10,
			Step: 0.1,
		},
	},
	PolarParams{
		Name:  "cardioid",
		Scale: 200,
		GetRadius: func(angle float64) float64 {
			return math.Sin(angle / 2)
		},
		Angle: AngleSets{
			Min:  0,
			Max:  math.Pi * 2,
			Step: 0.01,
		},
	},
	PolarParams{
		Name:  "lemniscate",
		Scale: 2,
		GetRadius: func(angle float64) float64 {
			const a = 100
			c := math.Cos(2 * angle)
			if c < 0 {
				return 0
			}
			return math.Sqrt(a * a * c)
		},
		Angle: AngleSets{
			Min:  0,
			Max:  math.Pi*2 + 0.1,
			Step: math.Pi / 100,
		},
	},
	PolarParams{
		Name:  "cannabis",
		Scale: 200,
		GetRadius: func(angle float64) float64 {

			a := 0.3
			return a * (1.0 + 9.0/10.0*math.Cos(8.0*angle)) * (1.0 + 1.0/10.0*math.Cos(24.0*angle)) * (9.0/10.0 + 1.0/10.0*math.Cos(200.0*angle)) * (1.0 + math.Sin(angle))
		},
		Angle: AngleSets{
			Min:  0,
			Max:  math.Pi * 2,
			Step: 0.01,
		},
	},
	PolarParams{
		Name:  "ellipse",
		Scale: 200,
		GetRadius: func(angle float64) float64 {

			var (
				a = 0.7
				e = 0.6
				l = a * (1 - e*e)
			)

			return l / (1.0 - e*math.Cos(angle))
		},
		Angle: AngleSets{
			Min:  0,
			Max:  math.Pi * 2,
			Step: 0.01,
		},
	},
	PolarParams{
		Name:  "custom",
		Scale: 200,
		GetRadius: func(angle float64) float64 {
			return math.Sin(angle * 3)
		},
		Angle: AngleSets{
			Min:  0,
			Max:  math.Pi * 2,
			Step: 0.01,
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

func Create(dir string, params PolarParams) {

	width, height := 512, 512

	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, width, height)
	canvas, _ := cairo.NewCanvas(surface)

	canvas.SetLineJoin(cairo.LINE_JOIN_ROUND)

	canvas.SetLineWidth(1.0)
	canvas.SetSourceRGB(0.5, 0.5, 0.5)
	DrawAxes(canvas, width, height)

	canvas.SetLineWidth(2)
	canvas.SetSourceRGB(0.5, 0, 0)
	PolarDraw(canvas, width, height, params)

	fileName := filepath.Join(dir, fmt.Sprintf("polar-%s.png", params.Name))

	surface.WriteToPNG(fileName)
	surface.Finish()
}

type PolarParams struct {
	Name      string
	Scale     float64
	GetRadius func(Angle float64) float64

	Angle AngleSets
}

type AngleSets struct {
	Min  float64
	Max  float64
	Step float64
}

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

func DrawAxes(canvas *cairo.Canvas, width, height int) {

	var (
		x0 = float64(width) * 0.5
		y0 = float64(height) * 0.5
	)

	rd := float64(40)

	k := 6

	m := 40

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

	n := 32

	du = 2 * math.Pi / float64(n)

	u := float64(0)
	for i := 0; i < n; i++ {

		s, c := math.Sincos(u)

		r := rd * float64(k)
		x := x0 + r*c
		y := y0 + r*s

		canvas.MoveTo(x0, y0)
		canvas.LineTo(x, y)

		u += du
	}

	canvas.Stroke()
}

func PolarDraw(canvas *cairo.Canvas, width, height int, params PolarParams) {

	var (
		x, y float64

		center = Decart{
			X: float64(width) * 0.5,
			Y: float64(height) * 0.5,
		}

		p = Polar{Phi: params.Angle.Min}
	)

	var step = func() bool {

		if p.Phi > params.Angle.Max {
			return false
		}

		p.R = params.GetRadius(p.Phi)
		d := PolarToDecart(p)
		p.Phi += params.Angle.Step

		temp := d.Scale(params.Scale)
		temp = center.Add(temp)

		x, y = temp.X, temp.Y

		return true
	}

	if step() {
		canvas.MoveTo(x, y)
		for step() {
			canvas.LineTo(x, y)
		}
	}

	canvas.Stroke()
	//canvas.Fill()
}

func main() {

	//CairoHelloWorld()

	path := os.Args[0]
	dir, _ := filepath.Split(path)
	dir = filepath.Join(dir, "test")

	err := makeDir(dir)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, p := range ps {
		Create(dir, p)
	}

	//Create(dir, ps[4])
}
