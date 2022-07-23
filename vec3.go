package main

type Vec struct {
	e [3]float64
}

func NewVector(x, y, z float64) *Vec {
	return &Vec{e: [3]float64{x, y, z}}
}

// Basic gettings.
func (v *Vec) X() float64 {
	return v.e[0]
}

func (v *Vec) Y() float64 {
	return v.e[1]
}

func (v *Vec) Z() float64 {
	return v.e[2]
}

func VecSub(v1, v2 Vec) Vec {
	return *NewVector(v1.X()-v2.X(), v1.Y()-v2.Y(), v1.Z()-v2.Z())
}

func VecDot(v1, v2 Vec) float64 {
	return v1.e[0]*v2.e[0] + v1.e[1]*v2.e[1] + v1.e[2]*v2.e[2]
}
