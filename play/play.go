package play

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/amyadzuki/amysgame/events"
	"github.com/amyadzuki/amysgame/game"
	"github.com/amyadzuki/amysgame/human"

	"github.com/amyadzuki/amystuff/gamecam"

	"github.com/amyadzuki/amygolib/states"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/loader/collada"
	"github.com/g3n/engine/loader/obj"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"github.com/g3n/engine/window"
)

var _ = fmt.Printf
var _ = collada.Decode
var _ = obj.Decode
var _ = core.NewNode
var _ = gls.NewVBO

var Fps, Ping int
var start = false

type MeType struct {
	*human.Human
}

func (m *MeType) FacingNormalized() (x float64, y float64) {
	y, x = math.Sincos(math.Pi * 0.5)
	return
}

func (m *MeType) Position() (position math32.Vector3) {
	return
}

var Me MeType
var State states.State

func init() {
	Me.Human = human.New()
}

func update() {
	age := float64(game.Main.CharaDesignerBodyAge.Value())
	gender := float64(game.Main.CharaDesignerBodyGender.Value())
	muscle := float64(game.Main.CharaDesignerBodyMuscle.Value())
	weight := float64(game.Main.CharaDesignerBodyWeight.Value())
	human.Builder.Update(age, gender, muscle, weight)
	Me.Update(age, gender, muscle, weight)
}

func Play(ofs int) {
	fmt.Println("Offset:", ofs)
	game.Main.Init("My Game")
	game.Main.StartUp("log.txt") // calls flag.Parse
	human.Init(game.Main.Rend)
	fmt.Printf("Base:        %f\n", Me.Base())
	fmt.Printf("HeightToCap: %f\n", Me.HeightToCap())
	fmt.Printf("HeightToEye: %f\n", Me.HeightToEye())
	control := gamecam.New(&Me, game.Main.Camera, game.Main.Win)
	control.Gui = game.Main.Gui
	control.SetDefaultToScreen(false)

	sky, err := graphic.NewSkybox(graphic.SkyboxData{
		DirAndPrefix: filepath.Join(human.Assets, "clds1-"),
		Extension:    "png",
		Suffixes: [6]string{
			"px", // +X // east
			"nx", // -X // west
			"py", // +Y // north
			"ny", // -Y // south
			"pz", // +Z // up
			"nz", // -Z // down
		},
	})
	if err != nil {
		panic(err)
	}
	game.Main.Scene.Add(sky)

	game.Main.AddWidgetCharaChanger("Log in")
	game.Main.AddWidgetHint("Welcome to My Game!")
	game.Main.AddWidgetClose("     X     ")
	game.Main.AddWidgetFullScreen("   +   ", "   o   ")
	game.Main.AddWidgetIconify("   _   ")
	game.Main.AddWidgetHelp("   ?   ")
	game.Main.AddWidgetPing()
	game.Main.AddWidgetFps()
	game.Main.InitWindows()
	update()
	game.Main.CharaDesignerBodyApply.Subscribe(gui.OnClick, func(evname string, event interface{}) {
		update()
	})
	game.Main.WidgetHelp.SetEnabled(false)
	game.Main.RecalcDocks()
	game.Main.Win.Subscribe(window.OnKeyDown, events.OnKeyboardKey)
	game.Main.Win.Subscribe(window.OnKeyUp, events.OnKeyboardKey)

	states.Note = game.Main.Logs.LoggerNote
	states.Debug = game.Main.Logs.LoggerDebug

	MainMenu := gui.NewPanel(0, 0)
	MainMenu.SetLayout(gui.NewDockLayout())
	{
		b := gui.NewButton("NEW CHARACTER")
		b.SetLayoutParams(&gui.DockLayoutParams{gui.DockTop})
		b.SetHeight(40)
		b.Subscribe(gui.OnClick, func(evname string, event interface{}) {
			State.SetNext("chara designer")
		})
		MainMenu.Add(b)
	}

	State.Init(game.Main.Win.ShouldClose).SetData(game.Main).SetFps(30).OnEnter("main menu", func(state *states.State) {

		control.SetEnabled(false)
		control.SetDefaultToScreen(true)
		game.Main.Gui.Add(MainMenu)
		siw, sih := game.Main.Win.Size()
		sw, sh := float64(siw), float64(sih)
		hh := 0.5 * sh
		pw, ph := math.Phi*hh, hh
		cx, cy := sw*0.5, sh*0.5
		px, py := cx-0.5*pw, cy-0.5*ph
		MainMenu.SetWidth(float32(pw))
		MainMenu.SetHeight(float32(ph))
		MainMenu.SetPosition(float32(px), float32(py))
		MainMenu.SetColor4(&math32.Color4{0, 0, 0, 0.5})

		game.Main.Camera.SetPositionVec(&math32.Vector3{0, 0, 0})
		game.Main.Camera.LookAt(&math32.Vector3{0, 0, 1.0 - 0.03125})

	}).OnLeave(func(state *states.State) {

		game.Main.Gui.Remove(MainMenu)

	}).OnFrame(func(state *states.State) {

		// Render the root GUI panel using the specified camera
		rendered, err := game.Main.Rend.Render(game.Main.Camera)
		if err != nil {
			panic(err)
		}
		game.Main.Wm.PollEvents()

		// Update window and checks for I/O events
		if rendered || true {
			game.Main.Win.SwapBuffers()
		}

	}).OnEnter("chara designer", func(state *states.State) {

		control.SetEnabled(true)
		control.SetDefaultToScreen(true)

		game.Main.Camera.SetPositionVec(&math32.Vector3{0, 2, 1})
		game.Main.Camera.LookAt(&math32.Vector3{0, 0, 1})

		game.Main.Scene.Add(Me)
		game.Main.WindowCharaDesignerOpen()

	}).OnLeave(func(state *states.State) {

		game.Main.WindowCharaDesignerClose()
		game.Main.Scene.Remove(Me)

	}).OnFrame(func(state *states.State) {

		Me.MatSkin.Udata.SkinDelta.X = game.Main.CharaDesignerSkinHue.Value()
		Me.MatSkin.Udata.SkinDelta.Y = game.Main.CharaDesignerSkinSat.Value()
		Me.MatSkin.Udata.SkinDelta.Z = game.Main.CharaDesignerSkinVal.Value()
		Me.MatSkin.Udata.SkinDelta.W = game.Main.CharaDesignerSkinTone.Value()

		Me.MatSkin.Udata.UwFabric.R = game.Main.CharaDesignerUwFabricRed.Value()
		Me.MatSkin.Udata.UwFabric.G = game.Main.CharaDesignerUwFabricGreen.Value()
		Me.MatSkin.Udata.UwFabric.B = game.Main.CharaDesignerUwFabricBlue.Value()
		Me.MatSkin.Udata.UwDetail.R = game.Main.CharaDesignerUwDetailRed.Value()
		Me.MatSkin.Udata.UwDetail.G = game.Main.CharaDesignerUwDetailGreen.Value()
		Me.MatSkin.Udata.UwDetail.B = game.Main.CharaDesignerUwDetailBlue.Value()
		Me.MatSkin.Udata.UwDetail.A = game.Main.CharaDesignerUwDetailAlpha.Value()
		Me.MatSkin.Udata.UwTrim.R = game.Main.CharaDesignerUwTrimRed.Value()
		Me.MatSkin.Udata.UwTrim.G = game.Main.CharaDesignerUwTrimGreen.Value()
		Me.MatSkin.Udata.UwTrim.B = game.Main.CharaDesignerUwTrimBlue.Value()

		Me.MatEyes.Udata.Color.R = game.Main.CharaDesignerEyeRed.Value()
		Me.MatEyes.Udata.Color.G = game.Main.CharaDesignerEyeGreen.Value()
		Me.MatEyes.Udata.Color.B = game.Main.CharaDesignerEyeBlue.Value()

		// Render the root GUI panel using the specified camera
		rendered, err := game.Main.Rend.Render(game.Main.Camera)
		if err != nil {
			panic(err)
		}
		game.Main.Wm.PollEvents()

		// Update window and checks for I/O events
		if rendered || true {
			game.Main.Win.SwapBuffers()
		}

	}).OnFrame("play", func(state *states.State) {

	}).SetNext("main menu").Run()
}
