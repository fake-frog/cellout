package main

import (
	"github.com/gdamore/tcell/v2"
)

type Cell struct {
	char string
	relX int // relative to tile x
	relY int // relative to tile y
}

type Displayable interface {
	SetTileStyleSelect()
	SetTileStyleNormal()
}

type Tile struct {
	Cells       []Cell
	Style       tcell.Style
	styleNormal tcell.Style
	styleSelect tcell.Style
	Content     string
	X           int
	Y           int
	H           int
	W           int
}

func NewTile(x int, y int, w int, h int, content string) *Tile {
	defaultNormalStyle := tcell.StyleDefault.Foreground(tcell.ColorPurple).Background(tcell.ColorBlack)
	defaultSelectStyle := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorPurple)

	newTile := Tile{
		Style:       defaultNormalStyle,
		styleNormal: defaultNormalStyle,
		styleSelect: defaultSelectStyle,
		X:           x,
		Y:           y,
		W:           w,
		H:           h,
	}

	newTile.Cells = setCells(w, h, content)
	return &newTile
}

func (tile *Tile) SetTileStyleSelect() {
	tile.Style = tile.styleSelect
}

func (tile *Tile) SetTileStyleNormal() {
	tile.Style = tile.styleNormal
}

/*
     W
*---------*
|         |
|    .    | H
|  (X,Y)  |
*---------*

x and y in middle

*/

func setCells(w int, h int, content string) []Cell {
	newCells := make([]Cell, 0, w*h)
	for i := 0; i < w*h; i++ {
		currX := i % w
		currY := i / w
		newCell := Cell{relX: currX, relY: currY}
		// corners
		if (currX == 0 && currY == 0) || (currX == w-1 && currY == 0) || (currX == 0 && currY == h-1) || (currX == w-1 && currY == h-1) {
			newCell.char = "█"
		}
		//top and bottom
		if (currX > 0 && currX < w-1) && (currY == 0 || currY == h-1) {
			newCell.char = "─"
		}
		//left right
		if (currY > 0 && currY < h-1) && (currX == 0 || currX == w-1) {
			newCell.char = "│"
		}
		newCells = append(newCells, newCell)
	}
	return newCells
}
