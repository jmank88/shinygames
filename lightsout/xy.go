package main

import "image"

// An XY represents a pixel in the board image.
// 0,0 1,0
// 0,1 1,1
type XY struct {
	x, y int
}

// An IJ represents a position on the board.
// 0,0 1,0
// 0,1 1,1
type IJ struct {
	i, j int
}

// IJ converts an XY position in the rectangle to a position on the board.
// Returns false if the point is not a board position.
func (xy XY) IJ(rect image.Rectangle, numSpaces int) (IJ, bool) {
	x, y := xy.x, xy.y
	spaceWidth := (rect.Max.X - rect.Min.X) / numSpaces
	spaceHeight := (rect.Max.Y - rect.Min.Y) / numSpaces
	i := (x - rect.Min.X) / spaceWidth
	j := (y - rect.Min.Y) / spaceHeight
	if 0 <= i && i <= numSpaces && 0 <= j && j <= numSpaces {
		return IJ{i, j}, true
	}
	return IJ{}, false
}
