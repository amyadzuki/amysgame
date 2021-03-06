// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.
package human

import (
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/texture"
)

var (
	Assets    string
	Builder   *Human
	Eyes      *texture.Texture2D
	HairLong  *texture.Texture2D
	SkinDark  *texture.Texture2D
	SkinLight *texture.Texture2D
	Underwear *texture.Texture2D
)

func Init(rend *renderer.Renderer) {
	rend.AddShader("HumanSkinVs", HumanSkinVs)
	rend.AddShader("HumanSkinFs", HumanSkinFs)
	rend.AddProgram("HumanSkin", "HumanSkinVs", "HumanSkinFs")
	rend.AddShader("HumanEyesVs", HumanEyesVs)
	rend.AddShader("HumanEyesFs", HumanEyesFs)
	rend.AddProgram("HumanEyes", "HumanEyesVs", "HumanEyesFs")
	rend.AddShader("HumanHairVs", HumanHairVs)
	rend.AddShader("HumanHairFs", HumanHairFs)
	rend.AddProgram("HumanHair", "HumanHairVs", "HumanHairFs")
}

func Load(path string) *texture.Texture2D {
	tex, err := texture.NewTexture2DFromImage(path)
	if err != nil {
		panic(err)
	}
	tex.SetFlipY(false)
	return tex
}

func TryLoad(path string) *texture.Texture2D {
	tex, err := texture.NewTexture2DFromImage(path)
	if err != nil {
		return nil
	}
	tex.SetFlipY(false)
	return tex
}

// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.
