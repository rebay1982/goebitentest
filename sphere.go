package main

import (
	"image/color"
	"math"
)

type Sphere struct {
	center Vec
	radius float64
	color  color.RGBA
}

func NewSphere(x, y, z, r float64) Sphere {
	return NewColoredSphere(x, y, z, r, color.RGBA{0xff, 0xff, 0xff, 0xff})
}

func NewColoredSphere(x, y, z, r float64, clr color.RGBA) Sphere {
	return Sphere{
		center: *NewVector(x, y, z),
		radius: r,
		color:  clr,
	}
}

func (s Sphere) GetColor() color.RGBA {
	return s.color
}

func (s Sphere) IntersectSphere(O, D Vec) (t1, t2 float64) {

	// See page 25.
	t1 = math.MaxFloat64
	t2 = math.MaxFloat64
	r := s.radius
	CO := VecSub(O, s.center)

	a := VecDot(D, D)
	b := 2 * VecDot(CO, D)
	c := VecDot(CO, CO) - r*r

	discriminant := b*b - 4*a*c
	if discriminant >= 0 {
		t1 = (-b + math.Sqrt(discriminant)) / (2 * a)
		t2 = (-b - math.Sqrt(discriminant)) / (2 * a)
	}

	return
}
