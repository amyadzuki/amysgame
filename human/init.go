// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.
package human

import (
	"path/filepath"

	"github.com/amyadzuki/amygolib/dirs"

	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/texture"
)

var (
	Assets string

	Dirs *dirs.Dirs
	Builder *Human

	Eyes      *texture.Texture2D
	SkinDark  *texture.Texture2D
	SkinLight *texture.Texture2D
	Underwear *texture.Texture2D
)

func Init(rend *renderer.Renderer) {
	rend.AddShader("HumanEyesVs", HumanEyesVs)
	rend.AddShader("HumanEyesFs", HumanEyesFs)
	rend.AddProgram("HumanEyes", "HumanEyesVs", "HumanEyesFs")
	rend.AddShader("HumanSkinVs", HumanSkinVs)
	rend.AddShader("HumanSkinFs", HumanSkinFs)
	rend.AddProgram("HumanSkin", "HumanSkinVs", "HumanSkinFs")
}

func init() {
	Dirs = dirs.New("Amy", "amysgame") // TODO: FIXME: I need to update this here and also in root.go

	Assets = filepath.Join(Dirs.ExeDir(), "assets")
	SkinDark = TryLoad(filepath.Join(Assets, "hsv01-v3.png"))
	SkinLight = TryLoad(filepath.Join(Assets, "hsv03-v3.png"))
	Eyes = TryLoad(filepath.Join(Assets, "eyes-v4.png"))
	Underwear = TryLoad(filepath.Join(Assets, "under-v2.png"))

	Builder = New()
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
