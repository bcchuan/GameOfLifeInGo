package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	screenWidth := 160
	screenHeight := 120
	maxInitLiveCells := int((screenWidth * screenHeight) / 8)

	g := NewGame(screenWidth, screenHeight, maxInitLiveCells)

	ebiten.SetWindowSize(screenWidth*4, screenHeight*4)
	ebiten.SetWindowTitle("Game of Life (Ebitengine Demo)")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
