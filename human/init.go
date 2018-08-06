package human

import (
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/texture"
)

var (
	MainBuilder *Builder

	Eyes      *texture.Texture2D
	SkinDark  *texture.Texture2D
	SkinLight *texture.Texture2D
	Underwear *texture.Texture2D
)

func Init(rend *renderer.Renderer, obj, mtl, darkSkin, lightSkin, underwear, eyes string) {
	rend.AddShader("HumanEyesVs", HumanEyesVs)
	rend.AddShader("HumanEyesFs", HumanEyesFs)
	rend.AddProgram("HumanEyes", "HumanEyesVs", "HumanEyesFs")
	rend.AddShader("HumanSkinVs", HumanSkinVs)
	rend.AddShader("HumanSkinFs", HumanSkinFs)
	rend.AddProgram("HumanSkin", "HumanSkinVs", "HumanSkinFs")

	SkinDark = Load(darkSkin)
	SkinLight = Load(lightSkin)
	Underwear = Load(underwear)
	Eyes = Load(eyes)

	dec, err := obj.Decode(obj, mtl)
	if err != nil {
		panic(err)
	}

	if MainBuilder, err = New(dec); err != nil {
		panic(err)
	}
}

func Load(path string) *texture.Texture2D {
	tex, err := texture.NewTexture2DFromImage(path)
	if err != nil {
		panic(err)
	}
	tex.SetFlipY(false)
	return tex
}
