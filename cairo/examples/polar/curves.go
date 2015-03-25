package main

import (
	"math"
)

type Curve interface {
	Name() string
	RadiusByAngle(Angle float64) float64
}

//------------------------------------------------------
type Spiral struct {
	a float64
}

func NewSpiral(a float64) Curve {
	return &Spiral{a}
}

func (this *Spiral) Name() string {
	return "spiral"
}

func (this *Spiral) RadiusByAngle(Angle float64) float64 {
	return this.a * Angle
}

//------------------------------------------------------
// r= a * (1 - cos(angle))
type Cardioid struct {
	a float64
}

func NewCardioid(a float64) Curve {
	return &Cardioid{a}
}

func (this *Cardioid) Name() string {
	return "cardioid"
}

func (this *Cardioid) RadiusByAngle(Angle float64) float64 {
	return this.a * (1.0 - math.Cos(Angle))
}

//------------------------------------------------------
type Lemniscate struct {
	a float64
}

func NewLemniscate(a float64) Curve {
	return &Lemniscate{a}
}

func (this *Lemniscate) Name() string {
	return "lemniscate"
}

func (this *Lemniscate) RadiusByAngle(Angle float64) float64 {

	c := math.Cos(2 * Angle)
	if c < 0.0 {
		return 0.0
	}
	return math.Sqrt(this.a * this.a * c)
}

//------------------------------------------------------
type Cannabis struct {
	a float64
}

func NewCannabis(a float64) Curve {
	return &Cannabis{a}
}

func (this *Cannabis) Name() string {
	return "cannabis"
}

func (this *Cannabis) RadiusByAngle(Angle float64) float64 {

	return this.a * (1.0 + 9.0/10.0*math.Cos(8.0*Angle)) *
		(1.0 + 1.0/10.0*math.Cos(24.0*Angle)) *
		(9.0/10.0 + 1.0/10.0*math.Cos(200.0*Angle)) *
		(1.0 + math.Sin(Angle))
}

//------------------------------------------------------
// r= a * sin(k * angle)
type Rose struct {
	name string
	a    float64
	k    float64
}

func NewRose(s string, a, k float64) Curve {
	return &Rose{s, a, k}
}

func (this *Rose) Name() string {
	return this.name
}

func (this *Rose) RadiusByAngle(Angle float64) float64 {
	return this.a * math.Sin(Angle*this.k)
}

//------------------------------------------------------
type Circle struct {
	a float64
}

func NewCircle(a float64) Curve {
	return &Circle{a}
}

func (this *Circle) Name() string {
	return "circle"
}

func (this *Circle) RadiusByAngle(Angle float64) float64 {
	return 2.0 * this.a * math.Sin(Angle)
}

//------------------------------------------------------
type Strofoid struct {
	b float64
}

func NewStrofoid(a float64) Curve {
	return &Strofoid{a}
}

func (this *Strofoid) Name() string {
	return "strofoid"
}

func (this *Strofoid) RadiusByAngle(Angle float64) float64 {

	sin, cos := math.Sincos(Angle)

	strofoid := this.b * (1.0 + cos) / sin

	return strofoid
}

//------------------------------------------------------
type StrofoidKnot struct {
	a, b float64
}

func NewStrofoidKnot(c float64) Curve {
	return &StrofoidKnot{
		a: c,
		b: 2.0 * c,
	}
}

func (this *StrofoidKnot) Name() string {
	return "strofoid-knot"
}

func (this *StrofoidKnot) RadiusByAngle(Angle float64) float64 {

	sin, cos := math.Sincos(Angle)

	circle := 2.0 * this.a * sin
	strofoid := this.b * (1.0 + cos) / sin
	knot := circle - strofoid

	return knot
}

//------------------------------------------------------
type ParabolaKnot struct {
	a, p float64
}

func NewParabolaKnot(c float64) Curve {
	return &ParabolaKnot{
		a: c,
		p: c / 4.0,
	}
}

func (this *ParabolaKnot) Name() string {
	return "parabola-knot"
}

func (this *ParabolaKnot) RadiusByAngle(Angle float64) float64 {

	sin, cos := math.Sincos(Angle)

	circle := 2.0 * this.a * sin
	parabola := 2.0 * this.p * sin / (cos * cos)
	knot := circle - parabola

	return knot
}
