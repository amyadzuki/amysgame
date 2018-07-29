package human

import (
	"github.com/g3n/engine/renderer"
)

func InitShaders(rend *renderer.Renderer) {
	rend.AddShader("HumanVs", HumanVs)
	rend.AddShader("HumanFs", HumanFs)
	rend.AddProgram("Human", "HumanVs", "HumanFs")
	// TODO: do these return error codes?
}

var HumanVs = `
`

var HumanFs = `
`
