package main

import (
	"github.com/gdamore/tcell/v2"
)

func setTileSelect(tile *Tile) {
	tile.SetTileStyleSelect()
}

func setTileNormal(tile *Tile) {
	tile.SetTileStyleNormal()
}

func main() {
	s, _ := tcell.NewScreen()
	s.Init()
	defer s.Fini()

	testOut := Cellout{}

	testOut.Cols = 12
	testOut.ColSize = 12

	testOut.PutTile(NewTile(10, 10, 120, 12, "something"))
	testOut.PutTile(NewTile(30, 20, 15, 15, "something"))
	testOut.PutTile(NewTile(10, 0, 20, 20, "something"))

	for {
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				return
			}
			if ev.Key() == tcell.KeyDown {
				testOut.SelectNextTile(setTileSelect, setTileNormal, setTileSelect, Down)
			}
			if ev.Key() == tcell.KeyUp {
				testOut.SelectNextTile(setTileSelect, setTileNormal, setTileSelect, Up)
			}
			if ev.Key() == tcell.KeyRight {
				testOut.SelectNextTile(setTileSelect, setTileNormal, setTileSelect, Right)
			}
			if ev.Key() == tcell.KeyLeft {
				testOut.SelectNextTile(setTileSelect, setTileNormal, setTileSelect, Left)
			}
		case *tcell.EventResize:
			s.Sync()
		}

		testOut.DrawTiles(s)
		s.Show()
	}

}
