package main

import (
	"fmt"

	"github.com/gitchander/go-examples/cairo"
)

func main() {
	if err := ExampleMSet(); err != nil {
		fmt.Println(err.Error())
	}
}

func ExampleMSet() error {

	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, 512, 512)
	defer surface.Destroy()

	canvas, _ := cairo.NewCanvas(surface)
	defer canvas.Destroy()

	var (
		width  = surface.GetWidth()
		height = surface.GetHeight()
		stride = surface.GetStride()
		err    error
	)

	n := surface.GetDataLength()
	bs := make([]byte, n)

	//canvas.SetSourceRGB(1, 1, 1)
	//canvas.Paint()

	if err = surface.GetData(bs); err != nil {
		return err
	}

	renderMSet(bs, width, height, stride)

	if err = surface.SetData(bs); err != nil {
		return err
	}

	surface.WriteToPNG("fractal.png")

	return nil
}

func renderMSet(bs []byte, width, height, stride int) {

	var (
		dx = 4.0 / float64(width)
		dy = 4.0 / float64(height)
	)

	var clBackground, clForeground, clResult ColorRGBA
	clForeground.Set(0.0, 0.0, 0.0, 1.0)

	n := 200

	y := -2.0
	for pY := 0; pY < height; pY++ {
		x := -2.0
		for pX := 0; pX < width; pX++ {

			clForeground.Alpha = calcAlphaSubpixel3x3(x, y, dx, dy, n)

			i := pX * 4
			clBackground.Decode(bs[i:])
			clResult = clForeground.Over(clBackground)
			clResult.Encode(bs[i:])

			x += dx
		}
		bs = bs[stride:]
		y += dy
	}
}

var subpixelShifts3x3 = []float64{
	-1.0 / 3.0,
	0.0,
	+1.0 / 3.0,
}

func calcAlphaSubpixel3x3(x0, y0 float64, dx, dy float64, n int) float64 {

	shift := subpixelShifts3x3
	m := len(shift)

	count := 0
	for iX := 0; iX < m; iX++ {
		for iY := 0; iY < m; iY++ {

			x := x0 + dx*shift[iX]
			y := y0 + dy*shift[iY]

			i := MandelbrotSet(x, y, n)
			if i >= n {
				count++
			}
		}
	}

	alpha := float64(count) / float64(m*m)

	return alpha
}

var subpixelShifts4x4 = []float64{
	-3.0 / 8.0,
	-1.0 / 8.0,
	+1.0 / 8.0,
	+3.0 / 8.0,
}

func calcAlphaSubpixel4x4(x0, y0 float64, dx, dy float64, n int) float64 {

	shift := subpixelShifts4x4
	m := len(shift)

	count := 0
	for iX := 0; iX < m; iX++ {
		for iY := 0; iY < m; iY++ {

			x := x0 + dx*shift[iX]
			y := y0 + dy*shift[iY]

			i := MandelbrotSet(x, y, n)
			if i >= n {
				count++
			}
		}
	}

	alpha := float64(count) / float64(m*m)

	return alpha
}
