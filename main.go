package cellout

import (
	"github.com/gdamore/tcell/v2"
)

func main() {
	s, _ := tcell.NewScreen()
	s.Init()
	defer s.Fini()

	testOut := Cellout{}

	for {
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				return
			}
			if ev.Key() == tcell.KeyDown {

			}
			if ev.Key() == tcell.KeyUp {

			}
			if ev.Key() == tcell.KeyRight {

			}
			if ev.Key() == tcell.KeyLeft {

			}
		case *tcell.EventResize:
			s.Sync()
		}

		testOut.DrawTiles(s)
		s.Show()
	}

}
