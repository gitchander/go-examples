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

func CanvasFill(canvas *cairo.Canvas, color ColorRGB) {

	surface := canvas.GetTarget()
	if surface == nil {
		return
	}

	canvas.Save()
	canvas.SetSourceRGB(color.Red, color.Green, color.Blue)
	canvas.Rectangle(0.0, 0.0, float64(surface.GetWidth()), float64(surface.GetHeight()))
	canvas.Fill()
	canvas.Restore()
}
