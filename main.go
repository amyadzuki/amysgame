// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.
package main

import (
	"github.com/amyadzuki/amysgame/run"
	"github.com/amyadzuki/amysgame/vars"

	"github.com/amy911/env911"
	"github.com/amy911/env911/app"
	"github.com/amy911/env911/config"
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

	config.BoolVar(&vars.Debug, "debug", false, "Write more log output")
	config.BoolVar(&vars.Trace, "debugextra", false, "Write significantly more log output")

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

// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.
