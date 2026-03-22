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
	Cells     []Cell
	TLeft     *Tile
	TRight    *Tile
	TUp       *Tile
	TDown     *Tile
	Style     tcell.Style
	X         int
	Y         int
	H         int
	W         int
	traversed bool
}

type Cellout struct {
	Tiles            []Tile
	CurrentTile      Tile
	currentTileIndex int
}

// func sortTilesByLoc(tiles []tile) []tile {

// }

func (cellout *Cell) SelectNextCellX(onSelect func(Tile), onUnselect func(Tile), onEnter func(Tile)) {

}

func (cellout *Cell) SelectNextCellY(onSelect func(Tile), onUnselect func(Tile), onEnter func(Tile)) {

}

func (cellout *Cell) SelectPrevCellX(onSelect func(Tile), onUnselect func(Tile), onEnter func(Tile)) {

}

func (cellout *Cell) SelectPrevCellY(onSelect func(Tile), onUnselect func(Tile), onEnter func(Tile)) {

}

func drawTiles(tiles []Tile, screen tcell.Screen) {
	for _, tile := range tiles {
		for _, cell := range tile.Cells {
			// rest of string, count
			screen.Put(tile.X+cell.relX, tile.Y+cell.relY, cell.char, tile.Style)
		}
	}
}

func drawRecuriveTiles(tile *Tile, screen tcell.Screen) {
	tile.traversed = true
	if tile.TLeft != nil && !tile.TLeft.traversed {
		drawRecuriveTiles(tile.TLeft, screen)
	}
	if tile.TRight != nil && !tile.TRight.traversed {
		drawRecuriveTiles(tile.TRight, screen)
	}
	if tile.TUp != nil && !tile.TUp.traversed {
		drawRecuriveTiles(tile.TUp, screen)
	}
	if tile.TDown != nil && !tile.TDown.traversed {
		drawRecuriveTiles(tile.TDown, screen)
	}

	for _, cell := range tile.Cells {
		screen.Put(tile.X+cell.relX, tile.Y+cell.relY, cell.char, tile.Style)
	}
}

func mockCell(char string, w int, h int) []Cell {
	cells := make([]Cell, 0, w*h)
	for i := 0; i < w*h; i++ {
		cell := Cell{
			char: char,
			relX: i % w,
			relY: i / w,
		}
		cells = append(cells, cell)
	}
	return cells
}

func mockTiles(char string, tw int, th int, cw int, ch int) []Tile {
	gap := 1
	style := tcell.StyleDefault.Foreground(tcell.ColorPurple).Background(tcell.ColorBlack)
	tiles := make([]Tile, 0, tw*th)
	for i := 0; i < tw*th; i++ {
		tile := Tile{
			Cells: mockCell(char, cw, ch),
			Style: style,
			X:     i % tw * (cw + gap),
			Y:     i / tw * (ch + gap),
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
		Tiles: mockTiles("/", 10, 10, 4, 2),
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
