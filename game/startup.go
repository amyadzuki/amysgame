package game

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/amyadzuki/amygolib/str"

	"github.com/amyadzuki/amystuff/logs"

	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/window"
)

func (game *Game) StartUp(logPath string) (err error) {
	flag_debug := CommandLine.Bool("debug", false,
		"Log debug info (may slightly slow the game)")
	flag_trace := CommandLine.Bool("debugextra", false,
		"Log trace info (may drastically slow the game)")
	flag_quiet := CommandLine.Bool("quiet", false,
		"Silence -info- messages from the console")
	flag_fullscreen := CommandLine.Bool("fullscreen", false,
		"Launch fullscreen")
	flag_geometry := CommandLine.String("geometry", strconv.Itoa(DflWidth)+"x"+strconv.Itoa(DflHeight),
		"Window geometry (H, WxH, or WxH+X+Y)")
	flag_wm := CommandLine.String("wm", "glfw",
		"Window manager (one of: \"glfw\")")

	func() {
		defer func() {
			if r := recover(); r != nil {
				CommandLine.Bool("x", false, "This help text")
				switch r {
				case flag.ErrHelp:
					os.Exit(0)
				default:
					os.Exit(2)
				}
			}
		}()
		CommandLine.Parse(os.Args[1:])
	}()

	info, debug, trace := !*flag_quiet, *flag_debug, *flag_trace
	game.InfoDebug = debug || trace
	game.InfoTrace = trace
	if game.Logs, err = logs.New(logPath, info, debug, trace); err != nil {
		return
	}

	game.Major("Created process and initialized logging for \"" + game.Title + "\"")

	w, h, x, y, n := DflWidth, DflHeight, 0, 0, 0
	n, err = fmt.Sscanf(*flag_geometry, "%dx%d+%d+%d", &w, &h, &x, &y)
	if n < 1 || n > 4 || (n == 4 && err != nil) {
		game.Warn("could not parse window geometry \"" + *flag_geometry + "\"")
	}
	if n == 1 {
		h = w
		w = h * 4 / 3
	}

	if h < 120 {
		game.Warn("height was capped at the minimum height of 120")
		h = 120
	}

	if w < 160 {
		game.Warn("width was capped at the minimum width of 160")
		w = 160
	}

	fs := *flag_fullscreen

	wm := str.Simp(*flag_wm)
	switch wm {
	case "glfw":
	default:
		game.Warn("unsupported window manager \"" + wm + "\" changed to \"glfw\"")
		wm = "glfw"
	}

	if game.Wm, err = window.Manager(wm); err != nil {
		game.Error("window.Manager")
		return
	}

	startupMessage :=
		"Launching \"" + game.Title + "\" " + strconv.Itoa(w) + "x" +
			strconv.Itoa(h) + " at (" + strconv.Itoa(x) + ", " +
			strconv.Itoa(y) + ") "
	if fs {
		startupMessage += "fullscreen"
	} else {
		startupMessage += "windowed"
	}
	startupMessage += " using \""
	startupMessage += wm
	startupMessage += "\""
	game.Info(startupMessage)

	if game.Win, err = game.Wm.CreateWindow(w, h, game.Title, fs); err != nil {
		game.Error("game.Wm.CreateWindow")
		return
	}

	// OpenGL functions must be executed in the same thread
	// where the context was created (by `CreateWindow`)
	runtime.LockOSThread()

	// Create the OpenGL state
	if game.Gs, err = gls.New(); err != nil {
		game.Error("gls.New")
		return
	}

	width, height := game.SizeRecalc()
	game.ViewportFull()
	aspect := float32(float64(width) / float64(height))
	game.Camera = camera.NewPerspective(65, aspect, 1.0/128.0, 1024.0)
	game.Scene.Add(game.Camera)
	game.Camera.SetUp(&math32.Vector3{0, 0, 1})

	game.LightAmbient = light.NewAmbient(&math32.Color{1, 1, 1}, 0.5)
	game.Scene.Add(game.LightAmbient)

	gui.SetStyleDefault(&styles.AmyDark)

	game.RealRoot = gui.NewRoot(game.Gs, game.Win)
	game.RealRoot.SetSize(float32(width), float32(height))
	game.RealRoot.SetLayout(gui.NewFillLayout(true, true))

	game.Gui = gui.NewPanel(float32(width), float32(height))
	game.Gui.SetLayout(gui.NewDockLayout())
	game.RealRoot.Add(game.Gui)

	game.Rend = renderer.NewRenderer(game.Gs)
	if err := game.Rend.AddDefaultShaders(); err != nil {
		panic(err)
	}
	game.Rend.SetScene(game.Scene)
	game.Rend.SetGui(game.RealRoot)

	game.Win.Subscribe(window.OnWindowSize, game.onWinCh)
	game.Win.Subscribe(window.OnKeyDown, game.onKeyboardKey)
	game.Win.Subscribe(window.OnKeyUp, game.onKeyboardKey)
	game.Win.Subscribe(window.OnMouseDown, game.onMouseButton)
	game.Win.Subscribe(window.OnMouseUp, game.onMouseButton)
	game.Win.Subscribe(window.OnCursor, game.onMouseCursor)

	/* if glfwWin, ok := game.Win.(*glfw.Window); ok {
		game.NkCtx = nk.NkPlatformInit(glfwWin, nk.PlatformInstallCallbacks)
		game.NkAtlas = nk.NewFontAtlas()
		nk.NkFontStashBegin(&game.NkAtlas)
		game.NkSans = nk.NkFontAtlasAddFromBytes(game.NkAtlas,
			MustAsset("assets/FreeSans.ttf"), 16, nil)
		nk.NkFontStashEnd()
		if game.NkSans != nil {
			nk.NkStyleSetFont(game.NkCtx, game.NkSans.Handle())
		}
	} */

	return
}