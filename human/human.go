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
	"github.com/g3n/engine/loader/obj"
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

func New(dec *obj.Decoder) (h *Human, err error) {
	h = new(Human)
	err = h.Init(dec)
	return
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

func (h *Human) Init(dec *obj.Decoder) (err error) {
	h.Lock() ; defer h.Unlock()
	h.age, h.gender, h.muscle, h.weight = 0.5, 0.125, 0.5, 0.5
	h.base, h.fOfEye, h.hToCap, h.hToEye = 0, -0.125, 1.5, 1.14
	h.BufIndEyes = math32.NewArrayU32(0, 0)
	h.BufIndSkin = math32.NewArrayU32(0, 0)
	h.BufPosEyes = math32.NewArrayF32(0, 0)
	h.BufPosSkin = math32.NewArrayF32(0, 0)
	h.BufUvsEyes = math32.NewArrayF32(0, 0)
	h.BufUvsSkin = math32.NewArrayF32(0, 0)
	h.GeomEyes = geometry.NewGeometry()
	h.GeomSkin = geometry.NewGeometry()
	h.GroupEyes = h.GeomEyes.AddGroup(h.BufIndEyes.Len(), 0, 0)
	h.GroupSkin = h.GeomSkin.AddGroup(h.BufIndSkin.Len(), 0, 0)
	h.GroupEyes.Count = h.BufPosEyes.Len()
	h.GroupSkin.Count = h.BufPosSkin.Len()
	h.GeomEyes.SetIndices(h.BufIndEyes)
	h.GeomSkin.SetIndices(h.BufIndSkin)
	h.VboPos = gls.NewVBO(h.BufPos).AddAttrib(gls.VertexPosition)
	h.GeomEyes.AddVBO(h.VboPos)
	h.GeomSkin.AddVBO(h.VboPos)
	h.GeomEyes.AddVBO(VboUvs)
	h.GeomSkin.AddVBO(VboUvs)
	h.MatEyes = material.NewStandard(&math32.Color{1.0/3, 2.0/3, 1})
	h.MatEyes = new(EyesMaterial)
	h.MatEyes.Init()
	h.MatEyes.Udata.Color = math32.Color4{1.0/3, 2.0/3, 1, 1}
	h.MatEyes.AddTexture(Eyes)
	h.MatSkin = material.NewStandard(&math32.Color{1, 1, 1})
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
	return
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

var HumanInit func(*Human)
var HumanUpdate func(*Human, bool)

/*
func (human *Embed) Init(
	dec       *obj.Decoder,
	skinDark  *texture.Texture2D,
	skinLight *texture.Texture2D,
	skinDelta *math32.Vector4,
	eyes      *texture.Texture2D,
	eyeColor  *math32.Color4,
	underwear *texture.Texture2D,
	uwFabric  *math32.Color4,
	uwDetail  *math32.Color4,
	uwTrim    *math32.Color4,
) (err error) {
	human.Node = core.NewNode()
	for idx := 0; idx < len(dec.Objects); idx++ {
		var vbo *gls.VBO
		var mesh *graphic.Mesh
		name := dec.Objects[idx].Name
		switch name {
		case "high-poly.obj":
			var m *EyesMaterial
			vbo, m, mesh, err = NewMeshEyes(dec, eyes, eyeColor, &dec.Objects[idx])
			if vbo == nil {
				panic("VboEyes was nil")
			}
			human.VboEyes, human.MatEyes, human.Eyes = vbo, m, mesh
			_, highest, frontest, ok := ofsRange(dec, &dec.Objects[idx])
			if ! ok {
				fmt.Printf("No vertices in \"%s\"\n", name)
			}
			human.heightToEye = float64(highest) - HalfEyeHeight
			human.frontOfEye = float64(frontest)
		case "female_generic.obj":
			fallthrough
		case "male_generic.obj":
			var m *SkinMaterial
			vbo, m, mesh, err = NewMeshSkin(dec, skinDark, skinLight, skinDelta,
				underwear, uwFabric, uwDetail, uwTrim, &dec.Objects[idx])
			if vbo == nil {
				panic("VboSkin was nil")
			}
			human.VboSkin, human.MatSkin, human.Skin = vbo, m, mesh
			lowest, highest, _, ok := ofsRange(dec, &dec.Objects[idx])
			if ! ok {
				fmt.Printf("No vertices in \"%s\"\n", name)
			}
			human.base = float64(lowest)
			human.heightToCap = float64(highest) - human.base
		default:
			fmt.Printf("Skipping: \"%s\"\n", name)
			continue;
		}
		if err != nil {
			return
		}
		human.Node.Add(mesh)
	}
	return nil
}

func ofsRange(dec *obj.Decoder, object *obj.Object) (lowest, highest, frontest float32, ok bool) {
	for _, face := range object.Faces {
		for _, vertex := range face.Vertices {
			y := dec.Vertices[vertex + 1]
			z := dec.Vertices[vertex + 2]
			if !false || z >= 1 {
				if ok {
					if z < lowest {
						lowest = z
					}
					if z > highest {
						highest = z
					}
					if !Backest {
						if y < frontest {
							frontest = y
						}
					} else {
						if y > frontest {
							frontest = y
						}
					}
				} else {
					lowest = z
					highest = z
					frontest = y
					ok = true
				}
			}
		}
	}
	// fmt.Printf("\"%s\": %f v^ %f (%v)  |  %f\n", object.Name, lowest, highest, 0.5 * (float64(highest) - float64(lowest)), frontest)
	return
}
*/

var HalfEyeHeight float64 = 0.013799965381622314
var Backest = false
var VboUvs *gls.VBO

func init() {
	VboUvs = gls.NewVBO(data.Coords[:]).AddAttrib(gls.VertexTexcoord)
}

// COPYRIGHT © 2018 amyadzuki <amyadzuki@gmail.com> ALL RIGHTS RESERVED.
