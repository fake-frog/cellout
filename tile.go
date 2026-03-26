package cellout

import (
	"github.com/gdamore/tcell/v2"
)

type Displayable interface {
	SetTileStyleSelect()
	SetTileStyleNormal()
}

type Cell struct {
	char string
	relX int // relative to tile x
	relY int // relative to tile y
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

func MakeTile(x int, y int, w int, h int, content string) *Tile {

}

type Tile struct {
	Cells       []Cell
	Style       tcell.Style
	styleNormal tcell.Style
	styleSelect tcell.Style
	X           int
	Y           int
	H           int
	W           int
}
