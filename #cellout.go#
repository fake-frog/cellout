package cellout

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
	DrawTiles(tcell.Screen)
	getTileLoc(*Tile) (int, int)
	findNextTileIndex(tiles *[]Tile, dir Direction) int
	currTile() *Tile
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

func (cellout Cellout) DrawTiles(screen tcell.Screen) {
	for _, tile := range cellout.Tiles {
		for _, cell := range tile.Cells {
			// rest of string, count
			screen.Put(tile.X+cell.relX, tile.Y+cell.relY, cell.char, tile.Style)
		}
	}
}
