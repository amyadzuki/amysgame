// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.
package human

import (
	"math"
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
	yAtEye, zAtBot, zAtCap, zAtEye float64
	theta, _padding float64

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
	VboPosEyes *gls.VBO
	VboPosSkin *gls.VBO

	sync.RWMutex
	*core.Node

	handleVao uint32

	finalized bool
}

func New(gs *gls.GLS) *Human {
	return new(Human).Init(gs)
}

func (h *Human) FacingNormalized() (dx, dy float64) {
	// h.RLock(); defer h.RUnlock()
	dy, dx = math.Sincos(h.theta)
	return
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
	// h.RLock(); defer h.RUnlock()
	return h.finalized
}

func (h *Human) Init(gs *gls.GLS) *Human {
	h.Lock() ; defer h.Unlock()
	h.age, h.gender, h.muscle, h.weight = 0.5, 0.125, 0.5, 0.5
	// Calc'd later: yAtEye, zAtBot, zAtCap, zAtEye
	h.theta = -0.5*math.Pi
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
	//h.handleVao = gs.GenVertexArray()
	//h.GeomEyes.SetVAO(h.handleVao)
	//h.GeomSkin.SetVAO(h.handleVao)
	h.GroupEyes = h.GeomEyes.AddGroup(h.BufIndEyes.Len(), 0, 0)
	h.GroupSkin = h.GeomSkin.AddGroup(h.BufIndSkin.Len(), 0, 0)
	h.GroupEyes.Count = h.BufPos.Len()
	h.GroupSkin.Count = h.BufPos.Len()
	h.GeomEyes.SetIndices(h.BufIndEyes)
	h.GeomSkin.SetIndices(h.BufIndSkin)
	h.VboPosEyes = gls.NewVBO(h.BufPos).AddAttrib(gls.VertexPosition)
	h.VboPosSkin = gls.NewVBO(h.BufPos).AddAttrib(gls.VertexPosition)
	h.GeomEyes.AddVBO(h.VboPosEyes)
	h.GeomSkin.AddVBO(h.VboPosSkin)
	h.GeomEyes.AddVBO(VboUvsEyes)
	h.GeomSkin.AddVBO(VboUvsSkin)
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
	// h.RLock(); defer h.RUnlock()
	return h.age, h.gender, h.muscle, h.weight
}

func (h *Human) Position() (pos math32.Vector3) {
	// h.RLock(); defer h.RUnlock()
	return
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

func (h *Human) YAtEye() float64 {
	// h.RLock(); defer h.RUnlock()
	return h.yAtEye
}

func (h *Human) ZAtBot() float64 {
	// h.RLock(); defer h.RUnlock()
	return h.zAtBot
}

func (h *Human) ZAtCap() float64 {
	// h.RLock(); defer h.RUnlock()
	return h.zAtCap
}

func (h *Human) ZAtEye() float64 {
	// h.RLock(); defer h.RUnlock()
	return h.zAtEye
}

func (h *Human) update_unlocked(final bool) {
	if h.finalized {
		return
	}
	if HumanUpdate != nil {
		HumanUpdate(h, final)
	}
	{
		zMaxSkin, zMinSkin := 0.0, 0.0
		for _, index := range h.BufIndSkin {
			z := float64(h.BufPos[index+2])
			if z > zMaxSkin {
				zMaxSkin = z
			}
			if z < zMinSkin {
				zMinSkin = z
			}
		}
		h.zAtBot = zMinSkin
		h.zAtCap = zMaxSkin
	}
	{
		yMinEyes, zMaxEyes := 0.0, 0.0
		for _, index := range h.BufIndEyes {
			y, z := float64(h.BufPos[index+1]), float64(h.BufPos[index+2])
			if y < yMinEyes {
				yMinEyes = y
			}
			if z > zMaxEyes {
				zMaxEyes = z
			}
		}
		h.yAtEye = yMinEyes
		h.zAtEye = zMaxEyes - HalfEyeHeight
	}
}

var Backest = false
var HalfEyeHeight float64 = 0.013799965381622314
var HumanInit func(*Human)
var HumanUpdate func(*Human, bool)
var VboUvsEyes *gls.VBO
var VboUvsSkin *gls.VBO

func init() {
	VboUvsEyes = gls.NewVBO(data.Uvs[:]).AddAttrib(gls.VertexTexcoord)
	VboUvsSkin = gls.NewVBO(data.Uvs[:]).AddAttrib(gls.VertexTexcoord)
}

// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.
