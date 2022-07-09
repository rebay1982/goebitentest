package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

// Represents the game state
type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	return
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (width, height int) {

	return outsideWidth >> 1, outsideHeight >> 1
}

func main() {

	ebiten.SetWindowSize(screenWidth<<1, screenHeight<<1)
	ebiten.SetWindowTitle("goebitentest")
	ebiten.SetWindowResizable(false)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
