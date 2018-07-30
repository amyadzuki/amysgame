package human

import (
	"unsafe"

	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/material"
)

type HumanSkinMaterial struct {
	material.Standard
	uni gls.Uniform
	Udata HumanSkinMaterialUdata
}

type HumanSkinMaterialUdata struct {
	SkinColor [4]float32
	UwFabric  [4]float32
	UwDetail  [4]float32
	UwTrim    [4]float32
}

func (m *HumanSkinMaterial) RenderSetup(gs *gls.GLS) {
	m.Standard.RenderSetup(gs)
	location := m.uni.Location(gs)
	gs.Uniform4fvUP(location, int32(unsafe.Sizeof(m.Udata) / 16), unsafe.Pointer(&m.Udata))
}
