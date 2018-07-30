package human

import (
	"unsafe"

	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

type HumanSkinMaterial struct {
	material.Standard
	uni gls.Uniform
	Udata HumanSkinMaterialUdata
}

type HumanSkinMaterialUdata struct {
	SkinColor math32.Vector4
	UwFabric  math32.Color4
	UwDetail  math32.Color4
	UwTrim    math32.Color4
}

func (m *HumanSkinMaterial) Init() {
	m.Standard.Init("HumanSkin", &math32.Color{1, 0, 1})
	m.uni.Init("HumanSkin")
}

func (m *HumanSkinMaterial) RenderSetup(gs *gls.GLS) {
	m.Standard.RenderSetup(gs)
	location := m.uni.Location(gs)
	gs.Uniform4fvUP(location, int32(unsafe.Sizeof(m.Udata) / 16), unsafe.Pointer(&m.Udata))
}
