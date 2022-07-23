package main

import (
	"github.com/rebay1982/goevitentest"
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
func CanvasToViewport(cx, cy float64) (vpCoord Vec3) {
	vx := cx * vpWidth / canvasWidth
	vy := cy * vpHeight / canvasHeight
	vz = vpDist

	return Vec3{e: [3]float64{vx, vy, vz}}
}
