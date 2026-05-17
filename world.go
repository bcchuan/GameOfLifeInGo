package main

import (
	"math/rand/v2"
)

// World represents the game state.
type World struct {
	area   []bool
	width  int
	height int
}

// NewWorld creates a new world.
func NewWorld(width, height int, maxInitLiveCells int) *World {
	w := &World{
		area:   make([]bool, width*height),
		width:  width,
		height: height,
	}
	w.init(maxInitLiveCells)
	return w
}

// init initializes the world with a random state.
func (w *World) init(maxLiveCells int) {
	for i := 0; i < maxLiveCells; i++ {
		x := rand.IntN(w.width)
		y := rand.IntN(w.height)
		w.area[y*w.width+x] = true
	}
}

// Update game state by one tick.
func (w *World) Update() {
	width := w.width
	height := w.height
	next := make([]bool, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pop := neighbourCount(w.area, width, height, x, y)
			switch {
			case pop < 2:
				// rule 1. Any live cell with fewer than two live neighbours
				// dies, as if caused by under-population.
				next[y*width+x] = false

			case (pop == 2 || pop == 3) && w.area[y*width+x]:
				// rule 2. Any live cell with two or three live neighbours
				// lives on to the next generation.
				next[y*width+x] = true

			case pop > 3:
				// rule 3. Any live cell with more than three live neighbours
				// dies, as if by over-population.
				next[y*width+x] = false

			case pop == 3:
				// rule 4. Any dead cell with exactly three live neighbours
				// becomes a live cell, as if by reproduction.
				next[y*width+x] = true
			}
		}
	}
	w.area = next
}

// neighbourCount calculates the Moore neighborhood of (x, y).
//
// ┌──────┬─────────────────────────────────┐
// │ Step │          What happens           │
// ├──────┼─────────────────────────────────┤
// │ 1    │ Start counter at 0              │
// ├──────┼─────────────────────────────────┤
// │ 2    │ Visit all 8 surrounding squares │
// ├──────┼─────────────────────────────────┤
// │ 3    │ Skip the cell itself            │
// ├──────┼─────────────────────────────────┤
// │ 4    │ Skip squares off the grid edge  │
// ├──────┼─────────────────────────────────┤
// │ 5    │ Count each living neighbor      │
// ├──────┼─────────────────────────────────┤
// │ 6    │ Return the total                │
// └──────┴─────────────────────────────────┘
//
// Imagine a grid of squares, like graph paper. Each square can be either
// alive or dead. To decide if a cell lives or dies next turn, we need to
// count how many of its neighbors (the squares touching it) are alive.
//
// Every cell has up to 8 neighbors — the squares directly above, below, left,
// right, and the 4 diagonals. Think of it like the 8 squares surrounding a
// king in chess.

// We're given the grid a, its size (width, height), and the position (x, y)
// of the cell we're checking.
//
func neighbourCount(a []bool, width, height, x, y int) int {

// Start a counter at zero. We'll add 1 each time we find a living neighbor.
//
	c := 0
// We loop through a 3×3 area centered on (x, y). Think of i and j as small
// steps:
// - i moves us left (-1), stay (0), or right (+1)
// - j moves us up (-1), stay (0), or down (+1)
//
// This visits all 9 squares in the 3×3 box (including the center cell itself)
//
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {
// When both steps are zero, we're looking at the cell itself — skip it!
// We only want to count neighbors, not the cell we're checking.
//
			if i == 0 && j == 0 {
				continue
			}
// Calculate the neighbor's position. If it falls off the edge of the grid
// (e.g., a cell on the left wall has no neighbor to its left), skip it.
// The grid doesn't wrap around.
//
			x2 := x + i
			y2 := y + j
			if x2 < 0 || y2 < 0 || width <= x2 || height <= y2 {
				continue
			}
// If the neighbor is alive, add 1 to our counter.
// (y2*width+x2 converts the 2D position into a spot in the flat list.)
//
			if a[y2*width+x2] {
				c++
			}
		}
	}
// Return the final count — could be anywhere from 0 to 8.
//
	return c
}
