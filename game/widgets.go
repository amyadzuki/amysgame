package game

import (
	"time"

	"github.com/amyadzuki/amysgame/vars"

	"github.com/suite911/amy911/crap/widget"

	"github.com/g3n/engine/gui"
)

func (game *Game) AddWidgetCharaChanger(label string) {
	if game.DockTop == nil {
		game.AddDockTop()
	}
	game.LabelCharaChangerBlank = label
	game.WidgetCharaChanger = gui.NewButton(label)
	game.WidgetCharaChanger.SetLayoutParams(&gui.DockLayoutParams{gui.DockLeft})
	game.WidgetCharaChanger.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		game.State.SetNext("chara select")
	})
	game.addDockSize(game.DockTop, game.WidgetCharaChanger)
	game.DockTop.Add(game.WidgetCharaChanger)
}

func (game *Game) AddWidgetClose(label string) {
	if game.DockTop == nil {
		game.AddDockTop()
	}
	game.WidgetClose = gui.NewButton(label)
	game.WidgetClose.SetLayoutParams(&gui.DockLayoutParams{gui.DockRight})
	game.WidgetClose.SetStyles(&vars.StylesEx.CloseButton)
	game.WidgetClose.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		game.WidgetClose.SetStyles(&vars.StylesEx.ClosingButton)
		if game.SoftQuit() > 0 {
			game.Quit()
		}
		go func() {
			time.Sleep(time.Second)
			game.AskQuit--
			if game.AskQuit < 0 {
				game.AskQuit = 0
			}
			game.WidgetClose.SetStyles(&vars.StylesEx.CloseButton)
		}()
	})
	game.addDockSize(game.DockTop, game.WidgetClose)
	game.DockTop.Add(game.WidgetClose)
}

func (game *Game) AddWidgetFps() {
	game.AddWidgetPerformance(&game.WidgetFps, 999999, " fps  ")
}

func (game *Game) AddWidgetFullScreen(labelFullScreen, labelWindow string) {
	if game.DockTop == nil {
		game.AddDockTop()
	}
	game.LabelFullScreen = labelFullScreen
	game.LabelWindow = labelWindow
	label := labelFullScreen
	if game.FullScreen() {
		label = labelWindow
	}
	game.WidgetFullScreen = gui.NewButton(label)
	game.WidgetFullScreen.SetLayoutParams(&gui.DockLayoutParams{gui.DockRight})
	game.WidgetFullScreen.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		game.ToggleFullScreen()
	})
	game.addDockSize(game.DockTop, game.WidgetFullScreen)
	game.DockTop.Add(game.WidgetFullScreen)
}

func (game *Game) AddWidgetHelp(label string) {
	if game.DockTop == nil {
		game.AddDockTop()
	}
	game.WidgetHelp = gui.NewButton(label)
	game.WidgetHelp.SetLayoutParams(&gui.DockLayoutParams{gui.DockRight})
	game.WidgetHelp.SetStyles(&vars.StylesEx.HelpButton)
	game.WidgetHelp.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		wh := !game.WantHelp
		game.WantHelp = wh
		if wh {
			game.WidgetHelp.SetStyles(&vars.StylesEx.HelpingButton)
		} else {
			game.WidgetHelp.SetStyles(&vars.StylesEx.HelpButton)
		}
	})
	game.addDockSize(game.DockTop, game.WidgetHelp)
	game.DockTop.Add(game.WidgetHelp)
}

func (game *Game) AddWidgetHint(label string) {
	if game.DockTop == nil {
		game.AddDockTop()
	}
	game.WidgetHint.Init(label)
	game.WidgetHint.Panel.SetLayoutParams(&gui.DockLayoutParams{gui.DockLeft})
	game.DockTop.Add(game.WidgetHint.Panel)
}

func (game *Game) AddWidgetIconify(label string) {
	if game.DockTop == nil {
		game.AddDockTop()
	}
	game.WidgetIconify = gui.NewButton(label)
	game.WidgetIconify.SetLayoutParams(&gui.DockLayoutParams{gui.DockRight})
	game.WidgetIconify.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		// TODO
	})
	game.addDockSize(game.DockTop, game.WidgetIconify)
	game.DockTop.Add(game.WidgetIconify)
}

func (game *Game) AddWidgetPerformance(w *widget.Performance, large int, label string) {
	w.Init(large, label)
	w.Outer.SetLayoutParams(&gui.DockLayoutParams{gui.DockRight})
	game.DockTop.Add(w.Outer)
}

func (game *Game) AddWidgetPing() {
	game.AddWidgetPerformance(&game.WidgetPing, 999999, " ms  ")
}
