package main

import (
	"github.com/rebay1982/goebitentest"
)

type Sphere struct {
	center Vec3
	radius float64
}

func NewSphere(x, y, z, r float64) *Spere {
	return &Sphere{
		center: &Vec3{e: [3]float64{x, y, z}},
		radius: r,
	}
}
