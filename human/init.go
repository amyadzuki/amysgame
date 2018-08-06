package human

import (
	"github.com/g3n/engine/loader/obj"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/texture"
)

var (
	MainBuilder *Builder

	Eyes       *texture.Texture2D
	SkinDarkF  *texture.Texture2D
	SkinDarkM  *texture.Texture2D
	SkinLightF *texture.Texture2D
	SkinLightM *texture.Texture2D
	UnderwearF *texture.Texture2D
	UnderwearM *texture.Texture2D
)

func Init(rend *renderer.Renderer, objF, mtlF, darkSkinF, lightSkinF, underwearF,
	objM, mtlM, darkSkinM, lightSkinM, underwearM, eyes string,
) {
	rend.AddShader("HumanEyesVs", HumanEyesVs)
	rend.AddShader("HumanEyesFs", HumanEyesFs)
	rend.AddProgram("HumanEyes", "HumanEyesVs", "HumanEyesFs")
	rend.AddShader("HumanSkinVs", HumanSkinVs)
	rend.AddShader("HumanSkinFs", HumanSkinFs)
	rend.AddProgram("HumanSkin", "HumanSkinVs", "HumanSkinFs")

	SkinDarkF = Load(darkSkinF)
	SkinDarkM = Load(darkSkinM)
	SkinLightF = Load(lightSkinF)
	SkinLightM = Load(lightSkinM)
	UnderwearF = Load(underwearF)
	UnderwearM = Load(underwearM)
	Eyes = Load(eyes)

	decF, err := obj.Decode(objF, mtlF)
	if err != nil {
		panic(err)
	}
	decM, err := obj.Decode(objM, mtlM)
	if err != nil {
		panic(err)
	}

	if MainBuilder, err = New(decF, decM); err != nil {
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
