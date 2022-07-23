package main

import (
	"image/color"
)

const (
	canvasWidth  = 640
	canvasHeight = 480

	vpWidth  float64 = 1
	vpHeight float64 = 1
	vpDist   float64 = 1
)

// CanvasToViewport takes canvas coordinates (pixel coordinates) and
//	returns their position in viweport coordinates.
func CanvasToViewport(cx, cy int) Vec {

	cxf := float64(cx)
	cyf := float64(cy)

	vx := cxf * vpWidth / canvasWidth
	vy := cyf * vpHeight / canvasHeight
	vz := vpDist

	return *NewVector(vx, vy, vz)
}

func TraceRay(O, D Vec, tmin, tmax float64) color.RGBA {
	return color.RGBA{0, 0, 0, 0xff}
}
