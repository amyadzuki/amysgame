// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.
package human

import (
	"sync"

	"github.com/amy911/amy911/maths"

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

	sync.RWMutex
	*core.Node

	handleVao uint32

	finalized bool
}

func New() *Human {
	return new(Human).Init()
}

func (h *Human) Base() float64 {
	h.RLock(); defer h.RUnlock()
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
	h.RLock(); defer h.RUnlock()
	return h.finalized
}

func (h *Human) FrontOfEye() float64 {
	h.RLock(); defer h.RUnlock()
	return h.fOfEye
}

func (h *Human) HeightToCap() float64 {
	h.RLock(); defer h.RUnlock()
	return h.hToCap
}

func (h *Human) HeightToEye() float64 {
	h.RLock(); defer h.RUnlock()
	return h.hToEye
}

func (h *Human) Init() *Human {
	h.Lock() ; defer h.Unlock()
	h.age, h.gender, h.muscle, h.weight = 0.5, 0.125, 0.5, 0.5
	h.base, h.fOfEye, h.hToCap, h.hToEye = 0, -0.125, 1.5, 1.14
	h.BufIndEyes = math32.NewArrayU32(0, 0)
	for _, index := range data.Ins["high-poly.obj"] {
		h.BufIndEyes.Append(index)
	}
	h.BufIndSkin = math32.NewArrayU32(0, 0)
	for _, index := range data.Ins["female_generic.obj"] {
		h.BufIndSkin.Append(index)
	}
	h.BufPos = math32.NewArrayF32(0, 0)
	for _, remap := range data.Remap {
		h.BufPos.Append(data.Float[3*remap+0][40])
		h.BufPos.Append(data.Float[3*remap+1][40])
		h.BufPos.Append(data.Float[3*remap+2][40])
	}
	h.GeomEyes = geometry.NewGeometry()
	h.GeomSkin = geometry.NewGeometry()
	h.handleVao = h.GeomEyes.VAO()
	badvao := h.GeomSkin.VAO()
	h.GeomSkin.SetVAO(h.handleVao)
	_ = badvao // TODO: delete it
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
	h.RLock(); defer h.RUnlock()
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
	VboUvs = gls.NewVBO(data.Uvs[:]).AddAttrib(gls.VertexTexcoord)
}

// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.
