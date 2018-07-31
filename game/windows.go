package game

import (
	"github.com/amyadzuki/amystuff/styles"

	"github.com/g3n/engine/gui"
)

func (game *Game) InitWindows() {
	game.WindowCharaDesigner = game.newWindow(480, 720, "Character Designer")
	{
		const sliderWidth = 480 - 60
		tree := gui.NewTree(480, 720)
		presets := tree.AddNode("Presets")
		a := gui.NewRadioButton("A")
		b := gui.NewRadioButton("B")
		c := gui.NewRadioButton("C")
		d := gui.NewRadioButton("D")
		a.SetGroup("preset")
		b.SetGroup("preset")
		c.SetGroup("preset")
		d.SetGroup("preset")
		presets.Add(a)
		presets.Add(b)
		presets.Add(c)
		presets.Add(d)

		skin := tree.AddNode("Skin")

		game.CharaDesignerSkinTone = gui.NewHSlider(sliderWidth, 20)
		game.CharaDesignerSkinTone.SetText("Skin Tone")
		game.CharaDesignerSkinTone.SetValue(0.25)
		skin.Add(game.CharaDesignerSkinTone)

		game.CharaDesignerSkinHue = gui.NewHSlider(sliderWidth, 20)
		game.CharaDesignerSkinHue.SetText("Hue Adjustment")
		game.CharaDesignerSkinHue.SetValue(0.5)
		skin.Add(game.CharaDesignerSkinHue)

		game.CharaDesignerSkinSat = gui.NewHSlider(sliderWidth, 20)
		game.CharaDesignerSkinSat.SetText("Saturation Adjustment")
		game.CharaDesignerSkinSat.SetValue(0.5)
		skin.Add(game.CharaDesignerSkinSat)

		game.CharaDesignerSkinVal = gui.NewHSlider(sliderWidth, 20)
		game.CharaDesignerSkinVal.SetText("Value Adjustment")
		game.CharaDesignerSkinVal.SetValue(0.5)
		skin.Add(game.CharaDesignerSkinVal)

		eyes := tree.AddNode("Eyes")

		game.CharaDesignerEyeRed = gui.NewHSlider(sliderWidth, 20)
		game.CharaDesignerEyeRed.SetText("Red")
		game.CharaDesignerEyeRed.SetValue(1.0/3)
		eyes.Add(game.CharaDesignerEyeRed)

		game.CharaDesignerEyeGreen = gui.NewHSlider(sliderWidth, 20)
		game.CharaDesignerEyeGreen.SetText("Green")
		game.CharaDesignerEyeGreen.SetValue(2.0/3)
		eyes.Add(game.CharaDesignerEyeGreen)

		game.CharaDesignerEyeBlue = gui.NewHSlider(sliderWidth, 20)
		game.CharaDesignerEyeBlue.SetText("Blue")
		game.CharaDesignerEyeBlue.SetValue(1)
		eyes.Add(game.CharaDesignerEyeBlue)

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
	win.SetResizable(true)
	win.SetLayout(gui.NewFillLayout(true, true))
	win.SetColor4(&styles.AmyDarkWindowContent)
	return win
}
