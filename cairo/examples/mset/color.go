package main

import (
	"errors"
	"math"
)

const COLOR_RGBA_SIZE = 4

const (
	encodeFactor = 255.0
	decodeFactor = 1.0 / encodeFactor
)

func round(x float64) float64 {
	return math.Floor(x + 0.5)
}

type ColorRGBA struct {
	Red, Green, Blue, Alpha float64
}

func NewColorRGBA(r, g, b, a float64) *ColorRGBA {

	c := &ColorRGBA{
		Red:   r,
		Green: g,
		Blue:  b,
		Alpha: a,
	}

	c.Normalize()

	return c
}

func normChan(channel float64) float64 {

	if channel < 0.0 {
		channel = 0.0
	}

	if channel > 1.0 {
		channel = 1.0
	}

	return channel
}

func (this *ColorRGBA) Normalize() {

	this.Red = normChan(this.Red)
	this.Green = normChan(this.Green)
	this.Blue = normChan(this.Blue)
	this.Alpha = normChan(this.Alpha)
}

func (this *ColorRGBA) Get() (r, g, b, a float64) {

	r = this.Red
	g = this.Green
	b = this.Blue
	a = this.Alpha

	return
}

func (this *ColorRGBA) Set(r, g, b, a float64) {

	this.Red = r
	this.Green = g
	this.Blue = b
	this.Alpha = a

	this.Normalize()
}

func (this *ColorRGBA) Encode(bs []byte) error {

	if len(bs) < COLOR_RGBA_SIZE {
		return errors.New("ColorRGBA.Encode(): wrong data size")
	}

	this.Normalize()

	bs[0] = byte(round(this.Blue * encodeFactor))
	bs[1] = byte(round(this.Green * encodeFactor))
	bs[2] = byte(round(this.Red * encodeFactor))
	bs[3] = byte(round(this.Alpha * encodeFactor))

	return nil
}

func (this *ColorRGBA) Decode(bs []byte) error {

	if len(bs) < COLOR_RGBA_SIZE {
		return errors.New("ColorRGBA.Decode(): wrong data size")
	}

	this.Blue = float64(bs[0]) * decodeFactor
	this.Green = float64(bs[1]) * decodeFactor
	this.Red = float64(bs[2]) * decodeFactor
	this.Alpha = float64(bs[3]) * decodeFactor

	this.Normalize()

	return nil
}

func (front *ColorRGBA) Over(back ColorRGBA) ColorRGBA {

	alpha := front.Alpha + back.Alpha*(1.0-front.Alpha)
	invAlpha := 1.0 / alpha

	red := (front.Red*front.Alpha + back.Red*(1.0-front.Alpha)) * invAlpha
	green := (front.Green*front.Alpha + back.Green*(1.0-front.Alpha)) * invAlpha
	blue := (front.Blue*front.Alpha + back.Blue*(1.0-front.Alpha)) * invAlpha

	return ColorRGBA{red, green, blue, alpha}
}
