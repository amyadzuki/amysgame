// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.
package human

import (
	"unsafe"

	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

// Skin ///////////////////////////////////////////////////////////////////////

type SkinMaterial struct {
	material.Standard
	uni gls.Uniform
	Udata SkinMaterialUdata
}

type SkinMaterialUdata struct {
	SkinDelta math32.Vector4
	UwFabric  math32.Color4
	UwDetail  math32.Color4
	UwTrim    math32.Color4
}

func (m *SkinMaterial) Init() {
	m.Standard.Init("HumanSkin", &math32.Color{1, 0, 1})
	m.uni.Init("HumanSkin")
}

func (m *SkinMaterial) RenderSetup(gs *gls.GLS) {
	m.Standard.RenderSetup(gs)
	location := m.uni.Location(gs)
	gs.Uniform4fvUP(location, int32(unsafe.Sizeof(m.Udata) / 16), unsafe.Pointer(&m.Udata))
}

// Eyes ///////////////////////////////////////////////////////////////////////

type EyesMaterial struct {
	material.Standard
	uni gls.Uniform
	Udata EyesMaterialUdata
}

type EyesMaterialUdata struct {
	Color math32.Color4
}

func (m *EyesMaterial) Init() {
	m.Standard.Init("HumanEyes", &math32.Color{1, 0, 1})
	m.uni.Init("HumanEyes")
}

func (m *EyesMaterial) RenderSetup(gs *gls.GLS) {
	m.Standard.RenderSetup(gs)
	location := m.uni.Location(gs)
	gs.Uniform4fvUP(location, int32(unsafe.Sizeof(m.Udata) / 16), unsafe.Pointer(&m.Udata))
}

// Hair ///////////////////////////////////////////////////////////////////////

type HairMaterial struct {
	material.Standard
	uni gls.Uniform
	Udata HairMaterialUdata
}

type HairMaterialUdata struct {
	Color math32.Color4
}

func (m *HairMaterial) Init() {
	m.Standard.Init("HumanHair", &math32.Color{1, 0, 1})
	m.uni.Init("HumanHair")
}

func (m *HairMaterial) RenderSetup(gs *gls.GLS) {
	m.Standard.RenderSetup(gs)
	location := m.uni.Location(gs)
	gs.Uniform4fvUP(location, int32(unsafe.Sizeof(m.Udata) / 16), unsafe.Pointer(&m.Udata))
}

// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.
