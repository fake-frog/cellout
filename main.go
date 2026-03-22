package main

import (
	"github.com/gdamore/tcell/v2"
)

type CellGrid interface {
	SelectNextCellX(onSelect func(Tile), onUnselect func(Tile), onEnter func(Tile))
	SelectNextCellY(onSelect func(Tile), onUnselect func(Tile), onEnter func(Tile))
	SelectPrevCellX(onSelect func(Tile), onUnselect func(Tile), onEnter func(Tile))
	SelectPrevCellY(onSelect func(Tile), onUnselect func(Tile), onEnter func(Tile))
}

type Cell struct {
	char string
	relX int // relative to tile x
	relY int // relative to tile y
}

type Tile struct {
	Cells []Cell
	Style tcell.Style
	X     int
	Y     int
	H     int
	W     int
}

type Cellout struct {
	Tiles []Tile
}

func drawTiles(tiles []Tile, screen tcell.Screen) {
	for _, tile := range tiles {
		for _, cell := range tile.Cells {
			// rest of string, count
			screen.Put(tile.X+cell.relX, tile.Y+cell.relY, cell.char, tile.Style)
		}
	}
}

func mockCell(char string, w int, h int) []Cell {
	cells := make([]Cell, 0, w*h)
	for i := 0; i < w*h; i++ {
		y := i / w
		x := i % w
		cell := Cell{
			char: char,
			relX: x,
			relY: y,
		}
		cells = append(cells, cell)
	}
	return cells
}

func mockTiles(char string, tw int, th int, cw int, ch int) []Tile {
	gap := 1
	style := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)
	tiles := make([]Tile, 0, tw*th)
	for i := 0; i < tw*th; i++ {
		y := i / tw * (ch + gap)
		x := i % tw * (cw + gap)
		tile := Tile{
			Cells: mockCell(char, cw, ch),
			Style: style,
			X:     x,
			Y:     y,
			H:     ch,
			W:     cw,
		}
		tiles = append(tiles, tile)
	}
	return tiles
}

func main() {
	s, _ := tcell.NewScreen()
	s.Init()
	defer s.Fini()

	testOut := Cellout{
		Tiles: mockTiles("*", 10, 10, 4, 2),
	}

	for {
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				return
			}
		case *tcell.EventResize:
			s.Sync()
		}

		drawTiles(testOut.Tiles, s)
		s.Show()
	}

}
