package main

import (
	"image"
	"image/color"
)

const spaces = 5

type board struct {
	lights [spaces][spaces]bool
}

// The load method sets the lights for level.
func (b *board) load(level int) {
	for i, bytes := range levels[level-1] {
		var j uint
		for j = 0; j < spaces; j++ {
			b.lights[i][j] = ((0X80 >> j) & bytes) > 0
		}
	}
}

// The click method translate a click at x, y on m into a board position,
// and flips the clicked light, as well as surrounding lights.
func (b *board) click(m *image.RGBA, x, y int) {
	ij, ok := XY{x, y}.IJ(m.Rect, spaces)
	if !ok {
		return
	}

	i := ij.i
	j := ij.j

	b.flip(i, j)
	b.flip(i+1, j)
	b.flip(i-1, j)
	b.flip(i, j+1)
	b.flip(i, j-1)
}

// The allOff method returns true if all lights are off.
func (b *board) allOff() bool {
	for i := 0; i < spaces; i++ {
		for j := 0; j < spaces; j++ {
			if b.lights[i][j] {
				return false
			}
		}
	}
	return true
}

// The flip method flips a lights state.  Out of bounds positions are ignored.
func (b *board) flip(i, j int) {
	if 0 <= i && i < spaces && 0 <= j && j < spaces {
		b.lights[i][j] = !b.lights[i][j]
	}
}

// The draw method draws each light on the board to m.
func (b *board) draw(m *image.RGBA) {
	spaceWidth := (m.Rect.Max.X - m.Rect.Min.X) / spaces
	spaceHeight := (m.Rect.Max.Y - m.Rect.Min.Y) / spaces
	for i := 0; i < spaces; i++ {
		for j := 0; j < spaces; j++ {
			b.drawLight(m, i, j, spaceWidth, spaceHeight)
		}
	}
}

const inset = 10

// The drawLight method draws a light as position i, j to m, with the given width and height.
func (b *board) drawLight(m *image.RGBA, i, j, width, height int) {
	c := color.Black
	if b.lights[i][j] {
		c = color.White
	}

	offX, offY := i* width, j* height
	for x := inset; x < width - inset; x++ {
		for y := inset; y < height - inset; y++ {
			m.Set(offX + x, offY + y, c)
		}
	}
}
