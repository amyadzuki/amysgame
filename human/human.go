// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.
package human

import (
	"sync"

	"github.com/amyadzuki/amygolib/maths"

	"github.com/amyadzuki/amysgame/data"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/math32"
)

type Human struct {
	age, gender, muscle, weight  float64
	base, fOfEye, hToCap, hToEye float64

	BufIndEyes math32.ArrayU32
	BufIndSkin math32.ArrayU32
	BufPos     math32.ArrayF32
	GeomEyes   *geometry.Geometry
	GeomSkin   *geometry.Geometry
	GroupEyes  *geometry.Group
	GroupSkin  *geometry.Group
	MatEyes    *EyesMaterial
	MatSkin    *SkinMaterial
	MeshEyes   *graphic.Mesh
	MeshSkin   *graphic.Mesh
	VboPos     *gls.VBO

	sync.Mutex
	*core.Node

	finalized bool
}

func New() *Human {
	return new(Human).Init()
}

func (h *Human) Base() float64 {
	return h.base
}

func (h *Human) Finalize() *Human {
	h.Lock() ; defer h.Unlock()
	if !h.finalized {
		h.update_unlocked(true)
		h.finalized = true
	}
	return h
}

func (h *Human) Finalized() bool {
	return h.finalized
}

func (h *Human) FrontOfEye() float64 {
	return h.fOfEye
}

func (h *Human) HeightToCap() float64 {
	return h.hToCap
}

func (h *Human) HeightToEye() float64 {
	return h.hToEye
}

func (h *Human) Init() *Human {
	h.Lock() ; defer h.Unlock()
	h.age, h.gender, h.muscle, h.weight = 0.5, 0.125, 0.5, 0.5
	h.base, h.fOfEye, h.hToCap, h.hToEye = 0, -0.125, 1.5, 1.14
	h.BufIndEyes = math32.NewArrayU32(0, 0)
	for _, index := range data.Indices["high-poly.obj"] {
		h.BufIndEyes.Append(index)
	}
	h.BufIndSkin = math32.NewArrayU32(0, 0)
	for _, index := range data.Indices["female_generic.obj"] {
		h.BufIndSkin.Append(index)
	}
	h.BufPos = math32.NewArrayF32(0, 0)
	for _, set := range data.Float {
		h.BufPos.Append(set[40])
	}
	h.GeomEyes = geometry.NewGeometry()
	h.GeomSkin = geometry.NewGeometry()
	h.GroupEyes = h.GeomEyes.AddGroup(h.BufIndEyes.Len(), 0, 0)
	h.GroupSkin = h.GeomSkin.AddGroup(h.BufIndSkin.Len(), 0, 0)
	h.GroupEyes.Count = h.BufPos.Len()
	h.GroupSkin.Count = h.BufPos.Len()
	h.GeomEyes.SetIndices(h.BufIndEyes)
	h.GeomSkin.SetIndices(h.BufIndSkin)
	h.VboPos = gls.NewVBO(h.BufPos).AddAttrib(gls.VertexPosition)
	h.GeomEyes.AddVBO(h.VboPos)
	h.GeomSkin.AddVBO(h.VboPos)
	h.GeomEyes.AddVBO(VboUvs)
	h.GeomSkin.AddVBO(VboUvs)
	h.MatEyes = new(EyesMaterial)
	h.MatEyes.Init()
	h.MatEyes.Udata.Color = math32.Color4{1.0/3, 2.0/3, 1, 1}
	h.MatEyes.AddTexture(Eyes)
	h.MatSkin = new(SkinMaterial)
	h.MatSkin.Init()
	h.MatSkin.Udata.SkinDelta = math32.Vector4{0.5, 0.5, 0.5, 0.25}
	h.MatSkin.Udata.UwFabric = math32.Color4{1, 1, 1, 1}
	h.MatSkin.Udata.UwDetail = math32.Color4{0.875, 0.875, 0.875, 0.5}
	h.MatSkin.Udata.UwTrim = math32.Color4{0xff/255.0, 0xb6/255.0, 0xc1/255.0, 1}
	h.MatSkin.AddTexture(SkinDark)
	h.MatSkin.AddTexture(SkinLight)
	h.MatSkin.AddTexture(Underwear)
	h.MeshEyes = graphic.NewMesh(h.GeomEyes, h.MatEyes)
	h.MeshSkin = graphic.NewMesh(h.GeomSkin, h.MatSkin)
	h.Node = core.NewNode()
	h.Node.Add(h.MeshSkin)
	h.Node.Add(h.MeshEyes)
	h.finalized = false
	if HumanInit != nil {
		HumanInit(h)
	}
	if HumanUpdate != nil {
		HumanUpdate(h, false)
	}
	return h
}

func (h *Human) Params() (float64, float64, float64, float64) {
	return h.age, h.gender, h.muscle, h.weight
}

func (h *Human) Update(age, gender, muscle, weight float64) *Human {
	h.Lock() ; defer h.Unlock()
	if !h.finalized {
		h.age = maths.ClampFloat64(age, 0, 1)
		h.gender = maths.ClampFloat64(gender, 0, 1)
		h.muscle = maths.ClampFloat64(muscle, 0, 1)
		h.weight = maths.ClampFloat64(weight, 0, 1)
	}
	h.update_unlocked(false)
	return h
}

func (h *Human) update_unlocked(final bool) {
	if h.finalized {
		return
	}
	if HumanUpdate != nil {
		HumanUpdate(h, final)
	}
}

var Backest = false
var HalfEyeHeight float64 = 0.013799965381622314
var HumanInit func(*Human)
var HumanUpdate func(*Human, bool)
var VboUvs *gls.VBO

func init() {
	VboUvs = gls.NewVBO(data.Coords[:]).AddAttrib(gls.VertexTexcoord)
}

// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.
