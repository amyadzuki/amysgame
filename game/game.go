package game

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/amyadzuki/amygolib/logs"

	"github.com/amyadzuki/amystuff/widget"

	"github.com/amyadzuki/client/loader"

	//	"github.com/g3n/engine/audio"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/window"

	"github.com/golang-ui/nuklear/nk"

	"github.com/go-gl/glfw/v3.2/glfw"
)

// const DflWidth, DflHeight = 960, 720
const DflWidth, DflHeight = 1600, 900

var _ = glfw.CreateWindow

var Main Game
type Game struct {
	logs.Logs

	Camera *camera.Perspective // TODO: camera.ICamera
	Win    window.IWindow
	Wm     window.IWindowManager

	DockBotLeft  *gui.Panel
	DockBotRight *gui.Panel
	DockTop      *gui.Panel

	LabelCharaChangerBlank string
	LabelFullScreen        string
	LabelWindow            string
	Title                  string

	Frame    int64
	SecFrame int64 // frame at start of this second

	WidgetFps           widget.Performance
	WidgetPing          widget.Performance
	WidgetHint          widget.Small
	WidgetCharaChanger  *gui.Button
	WidgetClose         *gui.Button
	WidgetFullScreen    *gui.Button
	WidgetHelp          *gui.Button
	WidgetIconify       *gui.Button
	WindowCharaDesigner *gui.Window
	WindowInventory     *gui.Window

	CharaDesignerBodyAge       *gui.Slider
	CharaDesignerBodyGender    *gui.Slider
	CharaDesignerBodyMuscle    *gui.Slider
	CharaDesignerBodyWeight    *gui.Slider
	CharaDesignerBodyApply     *gui.Button
	CharaDesignerSkinTone      *gui.Slider
	CharaDesignerSkinHue       *gui.Slider
	CharaDesignerSkinSat       *gui.Slider
	CharaDesignerSkinVal       *gui.Slider
	CharaDesignerEyeRed        *gui.Slider
	CharaDesignerEyeGreen      *gui.Slider
	CharaDesignerEyeBlue       *gui.Slider
	CharaDesignerUwFabricRed   *gui.Slider
	CharaDesignerUwFabricGreen *gui.Slider
	CharaDesignerUwFabricBlue  *gui.Slider
	CharaDesignerUwDetailRed   *gui.Slider
	CharaDesignerUwDetailGreen *gui.Slider
	CharaDesignerUwDetailBlue  *gui.Slider
	CharaDesignerUwDetailAlpha *gui.Slider
	CharaDesignerUwTrimRed     *gui.Slider
	CharaDesignerUwTrimGreen   *gui.Slider
	CharaDesignerUwTrimBlue    *gui.Slider

	NkCtx   *nk.Context
	NkAtlas *nk.FontAtlas
	NkSans  *nk.Font

	Gs           *gls.GLS
	LightAmbient *light.Ambient
	Rend         *renderer.Renderer
	RealRoot     *gui.Root
	Gui          *gui.Panel
	Scene        *core.Node

	w, h int

	AskQuit   int8
	WantHelp  bool
	HaveAudio bool
	InfoDebug bool
	InfoTrace bool
	MusicHush bool
	MusicMute bool

	OpenCharaDesigner bool
	OpenInventory     bool
}

func New(title string) (game *Game) {
	game = new(Game)
	game.Init(title)
	return
}

func (game *Game) AddDockBotLeft() {
	game.DockBotLeft = gui.NewPanel(0, 0)
	game.DockBotLeft.SetLayout(gui.NewDockLayout())
	game.Gui.Add(game.DockBotLeft)
}

func (game *Game) AddDockBotRight() {
	game.DockBotRight = gui.NewPanel(0, 0)
	game.DockBotRight.SetLayout(gui.NewDockLayout())
	game.Gui.Add(game.DockBotRight)
}

func (game *Game) AddDockTop() {
	game.DockTop = gui.NewPanel(0, 0)
	game.DockTop.SetLayout(gui.NewDockLayout())
	game.DockTop.SetLayoutParams(&gui.DockLayoutParams{gui.DockTop})
	game.Gui.Add(game.DockTop)
}

func (game *Game) Close() error {
	game.Dispose()
	return nil
}

func (game *Game) Dispose() {
	nk.NkPlatformShutdown()
}

func (game *Game) FullScreen() bool {
	return game.Win.FullScreen()
}

func (game *Game) Init(title string) {
	game.Title = title
	game.Scene = core.NewNode()
	game.Maker = loader.NewMaker()
	return
}

func (game *Game) Quit() {
	game.Win.SetShouldClose(true)
}

func (game *Game) RecalcDocks() {
	w, h := game.Size()
	w64, h64 := float64(w), float64(h)
	/*
		if game.DockTop != nil {
			game.DockTop.SetWidth(0)
			if game.WidgetCharaChanger != nil {
				game.addDockSize(game.DockTopLeft, game.WidgetCharaChanger)
			}
			if game.WidgetHint != nil {
				game.addDockSize(game.DockTopLeft, game.WidgetHint)
			}
		}
		if game.DockTopRight != nil {
			x := float32(w64 - float64(game.DockTopRight.TotalWidth()))
			game.DockTopRight.SetPosition(x, 0)
		}
	*/
	game.DockTop.SetWidth(float32(w))
	game.DockTop.SetHeight(float32(40))
	if game.DockBotLeft != nil {
		y := float32(h64 - float64(game.DockBotLeft.TotalHeight()))
		game.DockBotLeft.SetPosition(0, y)
	}
	if game.DockBotRight != nil {
		x := float32(w64 - float64(game.DockBotRight.TotalWidth()))
		y := float32(h64 - float64(game.DockBotRight.TotalHeight()))
		game.DockBotRight.SetPosition(x, y)
	}
}

func (game *Game) SetCharaDesigner(open bool) {
	if open == game.OpenCharaDesigner {
		return
	}
	game.OpenCharaDesigner = open
	if open {
		game.Gui.Add(game.WindowCharaDesigner)
	} else {
		game.Gui.Remove(game.WindowCharaDesigner)
	}
}

func (game *Game) SetFullScreen(fullScreen bool) {
	game.Win.SetFullScreen(fullScreen)
	if game.WidgetFullScreen != nil {
		label := game.LabelFullScreen
		if fullScreen {
			label = game.LabelWindow
		}
		game.WidgetFullScreen.Label.SetText(label)
	}
	game.onWinCh("", nil)
}

func (game *Game) SetHint(label string) {
	if game.WidgetHint.Label != nil {
		game.WidgetHint.Label.SetText(label)
		game.WidgetHint.Panel.SetWidth(game.WidgetHint.Label.TotalWidth())
		return
	}
	game.AddWidgetHint(label)
}

func (game *Game) SetInventory(open bool) {
	if open == game.OpenInventory {
		return
	}
	game.OpenInventory = open
	if open {
		game.Gui.Add(game.WindowInventory)
	} else {
		game.Gui.Remove(game.WindowInventory)
	}
}

func (game *Game) Size() (w, h int) {
	w, h = game.w, game.h
	return
}

func (game *Game) SizeRecalc() (w, h int) {
	w, h = game.Win.Size()
	game.w, game.h = w, h
	return
}

func (game *Game) SoftQuit() int8 {
	was := game.AskQuit
	now := int(was) + 1
	if now > 127 {
		now = 127
	}
	game.AskQuit = int8(now)
	return was
}

func (game *Game) ToggleCharaDesigner() {
	game.SetCharaDesigner(!game.OpenCharaDesigner)
}

func (game *Game) ToggleFullScreen() {
	game.SetFullScreen(!game.FullScreen())
}

func (game *Game) ToggleInventory() {
	game.SetInventory(!game.OpenInventory)
}

func (game *Game) ViewportFull() {
	w, h := game.Size()
	game.Gs.Viewport(0, 0, int32(w), int32(h))
	return
}

/*
func (game *Game) VolumeChanged() {
	if game.HaveAudio {
		loud := game.Settings.MusVolume.Value()
		if game.MusicHush {
			quiet := float32(float64(loud) * 0.5)
			game.PlayerMusic.SetGain(quiet)
		} else {
			game.PlayerMusic.SetGain(loud)
		}
	}
}
*/

// Internal functions

func (game *Game) addDockSize(p *gui.Panel, w gui.IPanel) {
	oldW, oldH := p.TotalWidth(), p.TotalHeight()
	newW, newH := oldW+w.TotalWidth(), w.TotalHeight()
	if oldH > newH {
		newH = oldH
	}
	p.SetWidth(newW)
	p.SetHeight(newH)
}

func printDefault(f *flag.Flag) {
	s := "  " + FlgHelpBeforeFlag + "-" + f.Name + FlgHelpAfterFlag
	name, usage := flag.UnquoteUsage(f)
	if len(name) > 0 {
		s += " " + FlgHelpBeforeType
		s += name // type name
		s += FlgHelpAfterType + " "
	}
	if len(s) <= 4 { // space, space, hyphen, character
		s += "\t"
	} else {
		s += ConsoleNewLineAndIndent
	}
	s += strings.Replace(usage, "\n", ConsoleNewLineAndIndent, -1)
	if !flag_isZeroValue(f, f.DefValue) {
		if reflect.TypeOf(f.Value).Elem().Kind() == reflect.String {
			s += fmt.Sprintf(" (default: %q)", f.DefValue)
		} else {
			s += fmt.Sprintf(" (default: %v)", f.DefValue)
		}
	}
	fmt.Fprint(CommandLine.Output(), s, "\n")
}

func flag_isZeroValue(flg *flag.Flag, value string) bool {
	typ := reflect.TypeOf(flg.Value)
	var z reflect.Value
	if typ.Kind() == reflect.Ptr {
		z = reflect.New(typ.Elem())
	} else {
		z = reflect.Zero(typ)
	}
	if value == z.Interface().(flag.Value).String() {
		return true
	}
	switch value {
	case "false", "", "0":
		return true
	}
	return false
}

var Usage = func() {
	fmt.Fprintf(CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	CommandLine.VisitAll(printDefault)
	fmt.Fprintln(CommandLine.Output(),
		"  "+FlgHelpBeforeFlag+"-h"+FlgHelpAfterFlag+" | "+
			FlgHelpBeforeFlag+"-help"+FlgHelpAfterFlag+" | "+
			FlgHelpBeforeFlag+"-?"+FlgHelpAfterFlag+
			ConsoleNewLineAndIndent+"Show this help message\n  "+
			"Note that you can use either one or two hyphens wherever one is shown")
}

var CommandLine = flag.NewFlagSet(os.Args[0], flag.PanicOnError)

func init() {
	CommandLine.Usage = Usage
}

const (
	ConsoleNewLineAndIndent = "\n      \t"
	VT100Bold               = "\x1b[1m"
	VT100Italic             = "\x1b[3m"
	VT100Underline          = "\x1b[4m"
	VT100Strike             = "\x1b[9m"
	VT100Reset              = "\x1b[0m\x1b[m"
	FlgHelpBeforeFlag       = VT100Bold
	FlgHelpAfterFlag        = VT100Reset
	FlgHelpBeforeType       = VT100Underline
	FlgHelpAfterType        = VT100Reset
)
