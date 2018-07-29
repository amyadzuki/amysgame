package game

import (
	"github.com/amyadzuki/amystuff/styles"

	"github.com/g3n/engine/gui"
)

func (game *Game) AddWindowInventory() {
	game.WindowInventory = gui.NewWindow(960, 720)
	game.WindowInventory.SetTitle("Inventory") // TODO: translate
	game.WindowInventory.SetPosition(60, 60)
	// Resizable windows are currently buggy.
	// game.WindowInventory.SetResizable(gui.ResizeAll)
	game.WindowInventory.SetLayout(gui.NewFillLayout(true, true))
	game.WindowInventory.SetColor4(&styles.AmyDarkWindowContent)
	game.Gui.Add(game.WindowInventory)
}
