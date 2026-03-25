package main

import (
	"github.com/gdamore/tcell/v2"
	"slices"
)

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

type CellGrid interface {
	SelectNextTile(onSelect func(*Tile), onUnselect func(*Tile), onEnter func(*Tile), dir Direction)
	PutTile(*Tile)
	PopTileAt(index int) *Tile
	getTileLoc(*Tile) (int, int)
	findNextTileIndex(tiles *[]Tile, dir Direction) int
	currTile() *Tile
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
	Tiles            []*Tile
	currentTileIndex int
	Cols             int
	ColSize          int
}

func clamp(v, min, max int) int {
	if v > max {
		return max
	}
	if v < min {
		return min
	}
	return v
}

// there is a vertual grid that we compare the position to TODO: better naming
func (cellout Cellout) getTileLoc(tile *Tile) (int, int) {
	X := (tile.X / cellout.ColSize) % cellout.Cols
	Y := (tile.Y / cellout.ColSize) * cellout.Cols
	return X, Y
}

func (cellout Cellout) currTile() *Tile {
	return cellout.Tiles[cellout.currentTileIndex]
}

func (cellout *Cellout) PutTile(tile *Tile) {
	newIndex := -1
	tile.X = clamp(tile.X, 0, cellout.Cols*cellout.ColSize) // theres probably a better way to handle this
	//	tileX, tileY := cellout.getTileLoc(tile)
	for i, currTile := range cellout.Tiles {
		//		currTileX, currTileY := cellout.getTileLoc(currTile)
		if tile.X+(tile.Y*cellout.Cols*cellout.ColSize) <= currTile.X+(currTile.Y*cellout.Cols*cellout.ColSize) {
			newIndex = i
			break
		}
	}
	if newIndex == -1 {
		cellout.Tiles = append(cellout.Tiles, tile)
		return
	}
	cellout.Tiles = slices.Insert(cellout.Tiles, newIndex, tile)
}

func isLeft(currTileLocX int, currTileLocY int, nextTileLocX int, nextTileLocY int) bool {
	return (nextTileLocX < currTileLocX &&
		nextTileLocY <= currTileLocY+2 &&
		nextTileLocY >= currTileLocY-2)
}

func isRight(currTileLocX int, currTileLocY int, nextTileLocX int, nextTileLocY int) bool {
	return (nextTileLocX > currTileLocX &&
		nextTileLocY <= currTileLocY+2 &&
		nextTileLocY >= currTileLocY-2)
}

func isUp(currTileLocX int, currTileLocY int, nextTileLocX int, nextTileLocY int) bool {
	return (nextTileLocY < currTileLocY &&
		nextTileLocX <= currTileLocX+2 &&
		nextTileLocX >= currTileLocX-2)
}

func isDown(currTileLocX int, currTileLocY int, nextTileLocX int, nextTileLocY int) bool {
	return (nextTileLocY > currTileLocY &&
		nextTileLocX <= currTileLocX+2 &&
		nextTileLocX >= currTileLocX-2)
}

func (cellout Cellout) findNextTileIndex(dir Direction) int {
	currTile := cellout.currTile()
	currTileLocX, currTileLocY := cellout.getTileLoc(currTile)

	switch dir {
	case Left:
		if cellout.currentTileIndex == 0 {
			return cellout.currentTileIndex
		}
		for i := cellout.currentTileIndex - 1; i >= 0; i-- {
			nextTileLocX, nextTileLocY := cellout.getTileLoc(cellout.Tiles[i])
			if isLeft(currTileLocX, currTileLocY, nextTileLocX, nextTileLocY) {
				return i
			}
		}
	case Right:
		if cellout.currentTileIndex == len(cellout.Tiles)-1 {
			return cellout.currentTileIndex
		}
		for i := cellout.currentTileIndex + 1; i < len(cellout.Tiles); i++ {
			nextTileLocX, nextTileLocY := cellout.getTileLoc(cellout.Tiles[i])
			if isRight(currTileLocX, currTileLocY, nextTileLocX, nextTileLocY) {
				return i
			}
		}
	case Up:
		if cellout.currentTileIndex == 0 {
			return cellout.currentTileIndex
		}
		for i := cellout.currentTileIndex - 1; i >= 0; i-- {
			nextTileLocX, nextTileLocY := cellout.getTileLoc(cellout.Tiles[i])
			if isUp(currTileLocX, currTileLocY, nextTileLocX, nextTileLocY) {
				return i
			}
		}
	case Down:
		if cellout.currentTileIndex == len(cellout.Tiles)-1 {
			return cellout.currentTileIndex
		}
		for i := cellout.currentTileIndex + 1; i < len(cellout.Tiles); i++ {
			nextTileLocX, nextTileLocY := cellout.getTileLoc(cellout.Tiles[i])
			if isDown(currTileLocX, currTileLocY, nextTileLocX, nextTileLocY) {
				return i
			}
		}
	}

	return cellout.currentTileIndex
}

func (cellout *Cellout) SelectNextTile(onSelect func(*Tile), onUnselect func(*Tile), onEnter func(*Tile), dir Direction) {

	//nextIndex := (cellout.currentTileIndex + 1) % len(cellout.Tiles)
	nextIndex := cellout.findNextTileIndex(dir)
	onUnselect(cellout.currTile())
	onSelect(cellout.Tiles[nextIndex])
	cellout.currentTileIndex = nextIndex
}

func drawTiles(tiles []*Tile, screen tcell.Screen) {
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
		cell := Cell{
			char: char,
			relX: i % w,
			relY: i / w,
		}
		cells = append(cells, cell)
	}
	return cells
}

func mockTiles(char string, tw int, th int, cw int, ch int) []*Tile {
	gap := 1
	style := tcell.StyleDefault.Foreground(tcell.ColorPurple).Background(tcell.ColorBlack)
	tiles := make([]*Tile, 0, tw*th)
	for i := 0; i < tw*th; i++ {
		tile := Tile{
			Cells: mockCell(char, cw, ch),
			Style: style,
			X:     i % tw * (cw + gap),
			Y:     i / tw * (ch + gap),
			H:     ch,
			W:     cw,
		}
		tiles = append(tiles, &tile)
	}
	return tiles
}

func changeTileStyleSelect(tile *Tile) {
	selectedStyle := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorPurple)
	tile.Style = selectedStyle
}

func changeTileStyleUnselect(tile *Tile) {
	selectedStyle := tcell.StyleDefault.Foreground(tcell.ColorPurple).Background(tcell.ColorBlack)
	tile.Style = selectedStyle
}

func main() {
	s, _ := tcell.NewScreen()
	s.Init()
	defer s.Fini()

	testOut := Cellout{}

	testOut.Cols = 12
	testOut.ColSize = 12

	testTile1 := Tile{
		Cells: mockCell("*", 4, 4),
		Style: tcell.StyleDefault.Foreground(tcell.ColorPurple).Background(tcell.ColorBlack),
		X:     50,
		Y:     45,
	}

	testTile2 := Tile{
		Cells: mockCell("/", 4, 4),
		Style: tcell.StyleDefault.Foreground(tcell.ColorPurple).Background(tcell.ColorBlack),
		X:     0,
		Y:     0,
	}

	testTile3 := Tile{
		Cells: mockCell("#", 4, 4),
		Style: tcell.StyleDefault.Foreground(tcell.ColorPurple).Background(tcell.ColorBlack),
		X:     12,
		Y:     0,
	}

	testTile4 := Tile{
		Cells: mockCell("%", 4, 4),
		Style: tcell.StyleDefault.Foreground(tcell.ColorPurple).Background(tcell.ColorBlack),
		X:     18,
		Y:     32,
	}

	testTile5 := Tile{
		Cells: mockCell("*", 4, 4),
		Style: tcell.StyleDefault.Foreground(tcell.ColorPurple).Background(tcell.ColorBlack),
		X:     0,
		Y:     13,
	}

	testTile6 := Tile{
		Cells: mockCell("/", 4, 4),
		Style: tcell.StyleDefault.Foreground(tcell.ColorPurple).Background(tcell.ColorBlack),
		X:     120,
		Y:     6,
	}

	testTile7 := Tile{
		Cells: mockCell("#", 4, 4),
		Style: tcell.StyleDefault.Foreground(tcell.ColorPurple).Background(tcell.ColorBlack),
		X:     24,
		Y:     19,
	}

	testTile8 := Tile{
		Cells: mockCell("%", 4, 4),
		Style: tcell.StyleDefault.Foreground(tcell.ColorPurple).Background(tcell.ColorBlack),
		X:     38,
		Y:     5,
	}

	testOut.PutTile(&testTile1)
	testOut.PutTile(&testTile2)
	testOut.PutTile(&testTile3)
	testOut.PutTile(&testTile4)
	testOut.PutTile(&testTile5)
	testOut.PutTile(&testTile6)
	testOut.PutTile(&testTile7)
	testOut.PutTile(&testTile8)

	for {
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				return
			}
			if ev.Key() == tcell.KeyDown {
				testOut.SelectNextTile(changeTileStyleSelect, changeTileStyleUnselect, changeTileStyleUnselect, Down)
			}
			if ev.Key() == tcell.KeyUp {
				testOut.SelectNextTile(changeTileStyleSelect, changeTileStyleUnselect, changeTileStyleUnselect, Up)
			}
			if ev.Key() == tcell.KeyRight {
				testOut.SelectNextTile(changeTileStyleSelect, changeTileStyleUnselect, changeTileStyleUnselect, Right)
			}
			if ev.Key() == tcell.KeyLeft {
				testOut.SelectNextTile(changeTileStyleSelect, changeTileStyleUnselect, changeTileStyleUnselect, Left)
			}
		case *tcell.EventResize:
			s.Sync()
		}

		drawTiles(testOut.Tiles, s)
		s.Show()
	}

}
