package game

import (
	"github.com/amyadzuki/amystuff/styles"

	"github.com/g3n/engine/gui"
)

func (game *Game) InitWindows() {
	game.WindowCharaDesigner = game.newWindow(480, 720, "Character Designer")
	{
		tree := gui.NewTree(480, 720)
		presets := tree.AddNode("Presets")
		presets.Add(gui.NewRadioButton("A").SetGroup("preset"))
		presets.Add(gui.NewRadioButton("B").SetGroup("preset"))
		presets.Add(gui.NewRadioButton("C").SetGroup("preset"))
		presets.Add(gui.NewRadioButton("D").SetGroup("preset"))
		skin := tree.AddNode("Skin")
		skin.Add(gui.NewHSlider(240, 20))
		skin.Add(gui.NewHSlider(240, 20))
		skin.Add(gui.NewHSlider(240, 20))
		skin.Add(gui.NewHSlider(240, 20))
		game.WindowCharaDesigner.Add(tree)
	}
	game.WindowInventory = game.newWindow(960, 720, "Inventory")
}

func (game *Game) WindowCharaDesignerClose() {
	game.Gui.Remove(game.WindowCharaDesigner)
}

func (game *Game) WindowCharaDesignerOpen() {
	game.Gui.Add(game.WindowCharaDesigner)
}

func (game *Game) WindowInventoryClose() {
	game.Gui.Remove(game.WindowInventory)
}

func (game *Game) WindowInventoryOpen() {
	game.Gui.Add(game.WindowInventory)
}

func (game *Game) newWindow(w, h float32, title string) *gui.Window {
	win := gui.NewWindow(w, h)
	win.SetTitle(title) // TODO: translate
	win.SetPosition(60, 60)
	// Resizable windows are currently buggy.
	// win.SetResizable(gui.ResizeAll)
	win.SetLayout(gui.NewFillLayout(true, true))
	win.SetColor4(&styles.AmyDarkWindowContent)
	return win
}
