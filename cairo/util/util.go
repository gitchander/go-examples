package util

import (
	"github.com/gitchander/go-examples/cairo"
)

type ColorRGB struct {
	Red, Green, Blue float64
}

func (c *ColorRGB) Normalize() {

	c.Red = normColorChannel(c.Red)
	c.Green = normColorChannel(c.Green)
	c.Blue = normColorChannel(c.Blue)
}

var normColorChannel = func(x float64) float64 {

	if x < 0.0 {
		x = 0.0
	}

	if x > 1.0 {
		x = 1.0
	}

	return x
}

type ColorRGBA struct {
	Red, Green, Blue, Alpha float64
}

func CanvasFillRGB(canvas *cairo.Canvas, color ColorRGB) {

	surface := canvas.GetTarget()
	if surface == nil {
		return
	}

	canvas.Save()
	canvas.SetSourceRGB(color.Red, color.Green, color.Blue)
	canvas.Paint()
	canvas.Restore()
}

func CanvasFillRGBA(canvas *cairo.Canvas, color ColorRGBA) {

	surface := canvas.GetTarget()
	if surface == nil {
		return
	}

	canvas.Save()
	canvas.SetSourceRGBA(color.Red, color.Green, color.Blue, color.Alpha)
	canvas.SetOperator(cairo.OPERATOR_SOURCE)
	canvas.Paint()
	canvas.Restore()
}
