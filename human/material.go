package human

import (
	"unsafe"

	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

// Skin ///////////////////////////////////////////////////////////////////////

type HumanSkinMaterial struct {
	material.Standard
	uni gls.Uniform
	Udata HumanSkinMaterialUdata
}

type HumanSkinMaterialUdata struct {
	SkinDelta math32.Vector4
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

// Eyes ///////////////////////////////////////////////////////////////////////

type HumanEyesMaterial struct {
	material.Standard
	uni gls.Uniform
	Udata HumanEyesMaterialUdata
}

type HumanEyesMaterialUdata struct {
	Color math32.Color4
}

func (m *HumanEyesMaterial) Init() {
	m.Standard.Init("HumanEyes", &math32.Color{1, 0, 1})
	m.uni.Init("HumanEyes")
}

func (m *HumanEyesMaterial) RenderSetup(gs *gls.GLS) {
	m.Standard.RenderSetup(gs)
	location := m.uni.Location(gs)
	gs.Uniform4fvUP(location, int32(unsafe.Sizeof(m.Udata) / 16), unsafe.Pointer(&m.Udata))
}
