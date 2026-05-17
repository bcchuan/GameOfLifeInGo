package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	world        *World
	pixels       []byte
}

func NewGame(screenWidth, screenHeight int, maxInitLiveCells int) *Game {
	g := &Game{
		world:        NewWorld(screenWidth, screenHeight, maxInitLiveCells),
		pixels:       make([]byte, screenWidth*screenHeight*4),
	}
	return g
}

// ebiten API
func (g *Game) Update() error {
	g.world.Update()
	return nil
}

// ebiten API
func (g *Game) Draw(screen *ebiten.Image) {
	g.paint(g.pixels)
	screen.WritePixels(g.pixels)
}

// A pixel takes 4 bytes: Red, Green, Blue, Alpha.
// A slice ([]byte) is a three-field struct:
//   a pointer to an underlying array, a length, and a capacity.
//
func (g *Game) paint(pix []byte) {
	for i, v := range g.world.area {
		if v {
			pix[4*i] = 0xff // Red
			pix[4*i+1] = 0xff // Green
			pix[4*i+2] = 0xff // Blue
			pix[4*i+3] = 0xff // Alpha
		} else {
			pix[4*i] = 0
			pix[4*i+1] = 0
			pix[4*i+2] = 0
			pix[4*i+3] = 0
		}
	}
}

// ebiten API
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.world.width, g.world.height
}

