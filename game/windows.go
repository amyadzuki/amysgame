package game

import (
	"github.com/amyadzuki/amystuff/styles"

	"github.com/g3n/engine/gui"
)

func (game *Game) InitWindows() {
	game.WindowCharaDesigner = game.newWindow(480, 720, "Character Designer")
	game.WindowCharaDesigner.SetCloseButton(false)
	{
		const sliderWidth = 480 - 60
		tree := gui.NewTree(480, 720)
		presets := tree.AddNode("Presets")
		presets.Add(gui.NewRadioButton("A").SetGroup("preset"))
		presets.Add(gui.NewRadioButton("B").SetGroup("preset"))
		presets.Add(gui.NewRadioButton("C").SetGroup("preset"))
		presets.Add(gui.NewRadioButton("D").SetGroup("preset"))

		body := tree.AddNode("Body")

		game.CharaDesignerBodyAge = gui.NewHSlider(sliderWidth, 20)
		game.CharaDesignerBodyAge.SetText("Apparent Age")
		game.CharaDesignerBodyAge.SetValue(0.5)
		body.Add(game.CharaDesignerBodyAge)

		game.CharaDesignerBodyGender = gui.NewHSlider(sliderWidth, 20)
		game.CharaDesignerBodyGender.SetText("Gender")
		game.CharaDesignerBodyGender.SetValue(0.125)
		body.Add(game.CharaDesignerBodyGender)

		game.CharaDesignerBodyMuscle = gui.NewHSlider(sliderWidth, 20)
		game.CharaDesignerBodyMuscle.SetText("Muscle")
		game.CharaDesignerBodyMuscle.SetValue(0.5)
		body.Add(game.CharaDesignerBodyMuscle)

		game.CharaDesignerBodyWeight = gui.NewHSlider(sliderWidth, 20)
		game.CharaDesignerBodyWeight.SetText("Weight")
		game.CharaDesignerBodyWeight.SetValue(0.5)
		body.Add(game.CharaDesignerBodyWeight)

		game.CharaDesignerBodyApply = gui.NewButton("Apply")
		game.CharaDesignerBodyApply.Subscribe(gui.OnClick, func(evname string, event interface{}) {
			game.UpdateMaker()
		})
		body.Add(game.CharaDesignerBodyApply)

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

		hair := tree.AddNode("Hair")
		_ = hair

		uw := tree.AddNode("Underwear")

		game.CharaDesignerUwFabricRed = gui.NewHSlider(sliderWidth, 20)
		game.CharaDesignerUwFabricRed.SetText("Fabric Red")
		game.CharaDesignerUwFabricRed.SetValue(1)
		uw.Add(game.CharaDesignerUwFabricRed)

		game.CharaDesignerUwFabricGreen = gui.NewHSlider(sliderWidth, 20)
		game.CharaDesignerUwFabricGreen.SetText("Fabric Green")
		game.CharaDesignerUwFabricGreen.SetValue(1)
		uw.Add(game.CharaDesignerUwFabricGreen)

		game.CharaDesignerUwFabricBlue = gui.NewHSlider(sliderWidth, 20)
		game.CharaDesignerUwFabricBlue.SetText("Fabric Blue")
		game.CharaDesignerUwFabricBlue.SetValue(1)
		uw.Add(game.CharaDesignerUwFabricBlue)

		game.CharaDesignerUwDetailRed = gui.NewHSlider(sliderWidth, 20)
		game.CharaDesignerUwDetailRed.SetText("Detail Red")
		game.CharaDesignerUwDetailRed.SetValue(0.875)
		uw.Add(game.CharaDesignerUwDetailRed)

		game.CharaDesignerUwDetailGreen = gui.NewHSlider(sliderWidth, 20)
		game.CharaDesignerUwDetailGreen.SetText("Detail Green")
		game.CharaDesignerUwDetailGreen.SetValue(0.875)
		uw.Add(game.CharaDesignerUwDetailGreen)

		game.CharaDesignerUwDetailBlue = gui.NewHSlider(sliderWidth, 20)
		game.CharaDesignerUwDetailBlue.SetText("Detail Blue")
		game.CharaDesignerUwDetailBlue.SetValue(0.875)
		uw.Add(game.CharaDesignerUwDetailBlue)

		game.CharaDesignerUwDetailAlpha = gui.NewHSlider(sliderWidth, 20)
		game.CharaDesignerUwDetailAlpha.SetText("Detail Opacity")
		game.CharaDesignerUwDetailAlpha.SetValue(0.5)
		uw.Add(game.CharaDesignerUwDetailAlpha)

		game.CharaDesignerUwTrimRed = gui.NewHSlider(sliderWidth, 20)
		game.CharaDesignerUwTrimRed.SetText("Trim Red")
		game.CharaDesignerUwTrimRed.SetValue(0xff/255.0)
		uw.Add(game.CharaDesignerUwTrimRed)

		game.CharaDesignerUwTrimGreen = gui.NewHSlider(sliderWidth, 20)
		game.CharaDesignerUwTrimGreen.SetText("Trim Green")
		game.CharaDesignerUwTrimGreen.SetValue(0xb6/255.0)
		uw.Add(game.CharaDesignerUwTrimGreen)

		game.CharaDesignerUwTrimBlue = gui.NewHSlider(sliderWidth, 20)
		game.CharaDesignerUwTrimBlue.SetText("Trim Blue")
		game.CharaDesignerUwTrimBlue.SetValue(0xc1/255.0)
		uw.Add(game.CharaDesignerUwTrimBlue)

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
