// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.
package main

import (
	"strconv"

	"github.com/amyadzuki/amysgame/data"
	"github.com/amyadzuki/amysgame/human"
	"github.com/amyadzuki/amysgame/run"
	"github.com/amyadzuki/amysgame/vars"

	"github.com/amy911/amy911/maths"
	"github.com/amy911/env911"
	"github.com/amy911/env911/config"

	"github.com/g3n/engine/math32"
)

const (
	DflWidth, DflHeight = 1600, 900
)

func init() {
	env911.InitAll("AMYSGAME_", nil, "amyadzuki", "amysgame") // TODO: better vendor and app names!!
}

func main() {
	copyright := config.Bool("copyright", false, "Print the copyright notice and exit")
	eula := config.Bool("eula", false, "Print the End User License Agreement (EULA) and exit")
	legal := config.Bool("legal", false, "Print the copyright notice and exit")
	license := config.Bool("license", false, "Print the End User License Agreement (EULA) and exit")
	version := config.Bool("version", false, "Print the version number and exit")
	quick := config.Bool("quick", false, "Skip the launcher and just play the game played previously")

	config.BoolVarP(&vars.Verbose, "verbose", "v", false, "Write more output")
	config.BoolVarP(&vars.Quiet, "quiet", "q", false, "Silence -info- messages from the console")

	config.BoolVar(&vars.FullScreen, "fullscreen", false, "Launch the game fullscreen")
	config.StringVar(&vars.Geometry, "geometry",
		strconv.Itoa(DflWidth)+"x"+strconv.Itoa(DflHeight),
		"Window geometry (H, WxH, or WxH+X+Y)")
	config.StringVar(&vars.WM, "wm", "glfw", "Window manager (one of: \"glfw\")")

	config.BoolVar(&vars.Debug, "debug", false, "Log debug info (may slightly slow the game)")
	config.BoolVar(&vars.Trace, "debugextra", false, "Log trace info (may drastically slow the game)")

	config.BoolVar(&vars.JSON, "json", false, "Use JSON for the data format")
	config.BoolVar(&vars.XML, "xml", false, "Use XML for the data format")
	config.BoolVar(&vars.YAML, "yaml", false, "Use YAML for the data format")

	config.LoadAndParse()

	want_copyright := *legal || *copyright
	want_license := *legal || *eula || *license
	switch {
	case want_copyright && want_license:
	case want_license:
	case want_copyright:
	case *version:
	default:
		if *quick {
			// TODO
		}
		run.Play()
	}
}

func init() {
	human.HumanInit = func(b *human.Human) {
		b.BufPos = math32.NewArrayF32(3*len(data.Remap), 3*len(data.Remap))
		*(b.VboPosEyes.Buffer()) = b.BufPos
		*(b.VboPosHair.Buffer()) = b.BufPos
		*(b.VboPosSkin.Buffer()) = b.BufPos
	}
	human.HumanUpdate = func(b *human.Human, final bool) {
		age, gender, muscle, weight := b.Params()
		var a0, a1, g0, g1, m0, m1, w0, w1 int
		switch {
		case age <= 0.25:
			a0, a1, age = 0, 1, age*4.0
		case age <= 0.5:
			a0, a1, age = 1, 2, maths.Fma(age, 4.0, -1.0)
		case age <= 0.75:
			a0, a1, age = 2, 3, maths.Fma(age, 4.0, -2.0)
		default:
			a0, a1, age = 3, 4, maths.Fma(age, 4.0, -3.0)
		}
		switch {
		case gender <= 0.5:
			g0, g1, gender = 0, 1, gender*2.0
		default:
			g0, g1, gender = 1, 2, maths.Fma(gender, 2.0, -1.0)
		}
		m0, m1, muscle = 0, 1, muscle
		switch {
		case weight <= 0.5:
			w0, w1, weight = 0, 1, weight*2.0
		default:
			w0, w1, weight = 1, 2, maths.Fma(weight, 2.0, -1.0)
		}

		a0 *= 18
		a1 *= 18
		g0 *= 6
		g1 *= 6
		m0 *= 3
		m1 *= 3
		w0 *= 1
		w1 *= 1

		buf := b.BufPos
		for idx, remap := range data.Remap {
			buf[3*idx+0] = updateCoord(data.Float[3*remap+0], age, gender, muscle, weight, a0, a1, g0, g1, m0, m1, w0, w1)
			buf[3*idx+1] = updateCoord(data.Float[3*remap+1], age, gender, muscle, weight, a0, a1, g0, g1, m0, m1, w0, w1)
			buf[3*idx+2] = updateCoord(data.Float[3*remap+2], age, gender, muscle, weight, a0, a1, g0, g1, m0, m1, w0, w1)
		}
		b.VboPosEyes.Update()
		b.VboPosHair.Update()
		b.VboPosSkin.Update()
	}
}

func updateCoord(data [90]float32, age, gender, muscle, weight float64, a0, a1, g0, g1, m0, m1, w0, w1 int) float32 {
	a0g0m0w := maths.Mix(float64(data[a0+g0+m0+w0]), float64(data[a0+g0+m0+w1]), weight)
	a0g0m1w := maths.Mix(float64(data[a0+g0+m1+w0]), float64(data[a0+g0+m1+w1]), weight)
	a0g0m__ := maths.Mix(a0g0m0w, a0g0m1w, muscle)
	a0g1m0w := maths.Mix(float64(data[a0+g1+m0+w0]), float64(data[a0+g1+m0+w1]), weight)
	a0g1m1w := maths.Mix(float64(data[a0+g1+m1+w0]), float64(data[a0+g1+m1+w1]), weight)
	a0g1m__ := maths.Mix(a0g1m0w, a0g1m1w, muscle)
	a0g____ := maths.Mix(a0g0m__, a0g1m__, gender)
	a1g0m0w := maths.Mix(float64(data[a1+g0+m0+w0]), float64(data[a1+g0+m0+w1]), weight)
	a1g0m1w := maths.Mix(float64(data[a1+g0+m1+w0]), float64(data[a1+g0+m1+w1]), weight)
	a1g0m__ := maths.Mix(a1g0m0w, a1g0m1w, muscle)
	a1g1m0w := maths.Mix(float64(data[a1+g1+m0+w0]), float64(data[a1+g1+m0+w1]), weight)
	a1g1m1w := maths.Mix(float64(data[a1+g1+m1+w0]), float64(data[a1+g1+m1+w1]), weight)
	a1g1m__ := maths.Mix(a1g1m0w, a1g1m1w, muscle)
	a1g____ := maths.Mix(a1g0m__, a1g1m__, gender)
	a______ := maths.Mix(a0g____, a1g____, age)
	return float32(a______)
}

// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.
