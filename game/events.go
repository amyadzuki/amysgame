package game

import (
	"github.com/amyadzuki/amysgame/events"
)

func init() {
	events.Map(events.ToggleFullScreen, func(id int, down bool, evname string, ev interface{}) {
		if !down {
			Main.ToggleFullScreen()
		}
	})
	events.Event[events.ToggleInventory] = func(id int, down bool, evname string, ev interface{}) {
		if !down {
			Main.ToggleInventory()
		}
	}
}

func (game *Game) onKeyboardKey(evname string, ev interface{}) {
}

func (game *Game) onMouseButton(evname string, ev interface{}) {
}

func (game *Game) onMouseCursor(evname string, ev interface{}) {
}

func (game *Game) onWinCh(evname string, ev interface{}) {
	if game.Win == nil {
		game.Warn("onWinCh but game.Win was nil")
		return
	}
	w, h := game.SizeRecalc()
	if game.Gs != nil {
		game.Gs.Viewport(0, 0, int32(w), int32(h))
	} else {
		game.Warn("onWinCh but game.GS was nil")
	}
	if game.RealRoot != nil {
		game.RealRoot.SetSize(float32(w), float32(h))
	} else {
		game.Warn("onWinCh but game.RealRoot was nil")
	}
	game.RecalcDocks()
	if game.Camera != nil {
		game.Camera.SetAspect(float32(float64(w) / float64(h)))
	} else {
		game.Warn("onWinCh but game.Camera was nil")
	}
}
