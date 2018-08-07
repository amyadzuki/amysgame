package human

import (
	"fmt"
	"sync"

	"github.com/amyadzuki/amygolib/maths"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/loader/obj"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
)

type Human struct {
	age, gender, muscle, weight  float64
	base, fOfEye, hToCap, hToEye float64

	GeomEyes  *geometry.Geometry
	GeomSkin  *geometry.Geometry
	GroupEyes *geometry.Group
	GroupSkin *geometry.Group
	MatEyes   *EyesMaterial
	MatSkin   *SkinMaterial
	MeshEyes  *graphic.Mesh
	MeshSkin  *graphic.Mesh
	VboEyes   *gls.VBO
	VboSkin   *gls.VBO

	*core.Node

	sync.Mutex

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
//	skin := &math32.Vector4{0.5, 0.5, 0.5, 0.25}
//	eyes := &math32.Color4{1.0/3.0, 2.0/3.0, 1, 1}
//	uwF := &math32.Color4{1, 1, 1, 1}
//	uwD := &math32.Color4{0.875, 0.875, 0.875, 0.5}
//	uwT := &math32.Color4{0xff/255.0, 0xb6/255.0, 0xc1/255.0, 1}
	h.age, h.gender, h.muscle, h.weight = 0.5, 0.125, 0.5, 0.5
	h.base, h.fOfEye, h.hToCap, h.hToEye = 0, -0.125, 1.5, 1.14
	h.GeomEyes = geometry.NewGeometry()
	h.GeomSkin = geometry.NewGeometry()
	h.GroupEyes = h.GeomEyes.AddGroup(h.IndicesEyes.Len(), 0, 0)
	h.GroupSkin = h.GeomSkin.AddGroup(h.IndicesSkin.Len(), 0, 0)
	h.GroupEyes.Count = h.PositionsEyes.Len()
	h.GroupSkin.Count = h.PositionsSkin.Len()
	h.GeomEyes.SetIndices(h.IndicesEyes)
	h.GeomSkin.SetIndices(h.IndicesSkin)
	h.VboEyes = gls.NewVBO(h.PositionsEyes).AddAttrib(gls.VertexPosition)
	h.VboSkin = gls.NewVBO(h.PositionsSkin).AddAttrib(gls.VertexPosition)
	h.GeomEyes.AddVBO(h.VboEyes)
	h.GeomSkin.AddVBO(h.VboSkin)
	h.GeomEyes.AddVBO(gls.NewVBO(h.UvsEyes).AddAttrib(gls.VertexTexcoord))
	h.GeomSkin.AddVBO(gls.NewVBO(h.UvsSkin).AddAttrib(gls.VertexTexcoord))
	h.MatEyes = material.NewStandard(&math32.Color{1.0/3, 2.0/3, 1})
	h.MatSkin = material.NewStandard(&math32.Color{1, 1, 1})
	h.MatEyes.AddTexture(Eyes)
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

var NewMeshEyes = func(
	dec       *obj.Decoder,
	eyes      *texture.Texture2D,
	color     *math32.Color4,
	object    *obj.Object,
) (*gls.VBO, *EyesMaterial, *graphic.Mesh, error) {
	geom, err := dec.NewGeometry(object)
	if err != nil {
		return nil, nil, nil, err
	}
	vbo := geom.VBO(gls.VertexPosition)
	mat := new(EyesMaterial)
	mat.Init()
	mat.Udata.Color = *color
	mat.AddTexture(eyes)
	mesh := graphic.NewMesh(geom, mat)
	return vbo, mat, mesh, nil
}

var NewMeshSkin = func(
	dec       *obj.Decoder,
	skinDark  *texture.Texture2D,
	skinLight *texture.Texture2D,
	skinDelta *math32.Vector4,
	underwear *texture.Texture2D,
	uwFabric  *math32.Color4,
	uwDetail  *math32.Color4,
	uwTrim    *math32.Color4,
	object    *obj.Object,
) (*gls.VBO, *SkinMaterial, *graphic.Mesh, error) {
	geom, err := dec.NewGeometry(object)
	if err != nil {
		return nil, nil, nil, err
	}
	vbo := geom.VBO(gls.VertexPosition)
	mat := new(SkinMaterial)
	mat.Init()
	mat.Udata.SkinDelta = *skinDelta
	mat.Udata.UwFabric = *uwFabric
	mat.Udata.UwDetail = *uwDetail
	mat.Udata.UwTrim = *uwTrim
	mat.AddTexture(skinDark)
	mat.AddTexture(skinLight)
	mat.AddTexture(underwear)
	mesh := graphic.NewMesh(geom, mat)
	return vbo, mat, mesh, nil
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

var HalfEyeHeight float64 = 0.013799965381622314
var Backest = false



