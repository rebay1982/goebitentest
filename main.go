package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

// Represents the game state
type Game struct {
	framebuffer *ebiten.Image
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.framebuffer.Fill(color.RGBA{255, 0, 0, 255})
	screen.DrawImage(g.framebuffer, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (width, height int) {

	return outsideWidth >> 1, outsideHeight >> 1
}

func main() {

	ebiten.SetWindowSize(screenWidth<<1, screenHeight<<1)
	ebiten.SetWindowTitle("goebitentest")
	ebiten.SetWindowResizable(false)

	game := &Game{framebuffer: ebiten.NewImage(screenWidth>>1, screenHeight>>1)}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
