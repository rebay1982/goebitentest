package main

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"

	"image/color"
	"log"
	"math"
)

const (
	screenWidth  = 1000
	screenHeight = 1000
)

// Represents the game state
type Game struct {
	framebuffer []byte
	scene       []Sphere
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//g.PutPixel(128, 128, red)

	screen.ReplacePixels(g.framebuffer)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (width, height int) {

	return screenWidth, screenHeight
}

// PutPixel writes a pixel in the frame buffer at position x, y
func (g *Game) PutPixel(x, y int, c color.RGBA) error {

	if x > (screenWidth) || x < 0 {
		return errors.New("Screen coord x out of bounds")
	}

	if y > (screenWidth) || y < 0 {
		return errors.New("Screen coord y out of bounds")
	}

	pxlPos := y*(screenWidth<<2) + x<<2
	R, G, B, _ := c.RGBA() // Ignore alpha channel.

	g.framebuffer[pxlPos+0] = (byte)(R & 0xff)
	g.framebuffer[pxlPos+1] = (byte)(G & 0xff)
	g.framebuffer[pxlPos+2] = (byte)(B & 0xff)
	g.framebuffer[pxlPos+3] = 0xff

	return nil
}

func main() {

	ebiten.SetWindowSize(screenWidth+5, screenHeight+5)
	ebiten.SetWindowTitle("goebitentest")
	ebiten.SetWindowResizable(false)

	game := &Game{
		framebuffer: make([]byte, (screenWidth*screenHeight)<<2),
		scene:       make([]Sphere, 3),
	}

	game.scene[0] = NewColoredSphere(0.0, -1.0, 3.0, 1.0, color.RGBA{0xff, 0x00, 0x00, 0xff})
	game.scene[1] = NewColoredSphere(1.0, 0.0, 4.0, 1.0, color.RGBA{0x00, 0x00, 0xff, 0xff})
	game.scene[2] = NewColoredSphere(-2.0, 0.0, 4.0, 1.0, color.RGBA{0xff, 0xff, 0x00, 0xff})

	drawScene(game)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func CanvasToPixelCoords(cX, cY int) (pX, pY int) {
	pX = cX + (screenWidth >> 1)
	pY = -cY + (screenHeight >> 1) - 1
	return
}

func drawScene(game *Game) {
	origin := *NewVector(0, 0, 0)

	cwHalf := canvasWidth >> 1
	chHalf := canvasHeight >> 1

	for x := -cwHalf; x < cwHalf; x++ {
		for y := -chHalf; y < chHalf; y++ {
			D := CanvasToViewport(x, y)
			color := TraceRay(origin, D, 1, math.MaxFloat64, game.scene)

			pX, pY := CanvasToPixelCoords(x, y)
			game.PutPixel(pX, pY, color)
		}
	}
}
