package human

import (
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/texture"
)

var (
	SkinDark  *texture.Texture2D // 683d2b
	SkinLight *texture.Texture2D // e2be9f
	Underwear *texture.Texture2D
)

var (
	SkinColorDark = math32.Color{0x68/255.0, 0x3d/255.0, 0x2b/255.0}
	SkinColorLight = math32.Color{0xe2/255.0, 0xbe/255.0, 0x9f/255.0}
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
