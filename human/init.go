package human

import (
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/texture"
)

var (
	SkinDark  *texture.Texture2D
	SkinLight *texture.Texture2D
	Eyes      *texture.Texture2D
	Underwear *texture.Texture2D
)

func Init(rend *renderer.Renderer, darkSkin, lightSkin, eyes, underwear string) {
	rend.AddShader("HumanSkinVs", HumanSkinVs)
	rend.AddShader("HumanSkinFs", HumanSkinFs)
	rend.AddProgram("HumanSkin", "HumanSkinVs", "HumanSkinFs")
	rend.AddShader("HumanEyesVs", HumanEyesVs)
	rend.AddShader("HumanEyesFs", HumanEyesFs)
	rend.AddProgram("HumanEyes", "HumanEyesVs", "HumanEyesFs")
	// TODO: do these return error codes?

	SkinDark = Load(darkSkin)
	SkinLight = Load(lightSkin)
	Eyes = Load(eyes)
	Underwear = Load(underwear)
}

func Load(path string) *texture.Texture2D {
	tex, err := texture.NewTexture2DFromImage(path)
	if err != nil {
		panic(err)
	}
	tex.SetFlipY(false)
	return tex
}
