package human

import (
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/texture"
)

var (
	SkinDark  *texture.Texture2D
	SkinLight *texture.Texture2D
	Underwear *texture.Texture2D
)

func Init(rend *renderer.Renderer) {
	rend.AddShader("HumanSkinVs", HumanSkinVs)
	rend.AddShader("HumanSkinFs", HumanSkinFs)
	rend.AddProgram("HumanSkin", "HumanSkinVs", "HumanSkinFs")
	// TODO: do these return error codes?

	SkinDark = Load("01-dark.png")
	SkinLight = Load("03-light.png")
	Underwear = Load("10-under.png")
}

func Load(path string) *texture.Texture2D {
	tex, err := texture.NewTexture2DFromImage(path)
	if err != nil {
		panic(err)
	}
	tex.SetFlipY(false)
	return tex
}
