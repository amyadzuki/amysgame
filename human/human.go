package human

import (
	"fmt"
	"strings"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/loader/obj"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
)

type Human struct {
	Node *core.Node

	Eyes *graphic.Mesh
	Skin *graphic.Mesh

	base        float64
	frontOfEye  float64
	heightToCap float64
	heightToEye float64
}

func New(
	dec       *obj.Decoder,
	skin      *texture.Texture2D,
	skinColor *math32.Color4,
	underwear *texture.Texture2D,
	uwFabric  *math32.Color4,
	uwDetail  *math32.Color4,
	uwTrim    *math32.Color4,
) (human *Human, err error) {
	human = new(Human)
	err = human.Init(dec, skin, skinColor, underwear, uwFabric, uwDetail, uwTrim)
	return
}

func (human *Human) Base() float64 {
	return human.base
}

func (human *Human) FrontOfEye() float64 {
	return human.frontOfEye
}

func (human *Human) HeightToCap() float64 {
	return human.heightToCap
}

func (human *Human) HeightToEye() float64 {
	return human.heightToEye
}

func (human *Human) Init(
	dec       *obj.Decoder,
	skin      *texture.Texture2D,
	skinColor *math32.Color4,
	underwear *texture.Texture2D,
	uwFabric  *math32.Color4,
	uwDetail  *math32.Color4,
	uwTrim    *math32.Color4,
) (err error) {
	human.Node = core.NewNode()
	for idx := 0; idx < len(dec.Objects); idx++ {
		var mesh *graphic.Mesh
		name := dec.Objects[idx].Name
		switch {
		case strings.HasSuffix(name, "-highpolyeyes"):
			mesh, err = dec.NewMesh(&dec.Objects[idx])
			human.Eyes = mesh
			_, highest, frontest, ok := ofsRange(dec, &dec.Objects[idx])
			if ! ok {
				fmt.Printf("No vertices in \"%s\"\n", name)
			}
			human.heightToEye = float64(highest) - HalfEyeHeight
			human.frontOfEye = float64(frontest)
		case strings.HasSuffix(name, "-female_generic"):
			fallthrough
		case strings.HasSuffix(name, "-male_generic"):
			mesh, err = NewMeshSkin(dec, skin, skinColor, underwear, uwFabric,
				uwDetail, uwTrim, &dec.Objects[idx])
			human.Skin = mesh
			lowest, highest, _, ok := ofsRange(dec, &dec.Objects[idx])
			if ! ok {
				fmt.Printf("No vertices in \"%s\"\n", name)
			}
			human.base = float64(lowest)
			human.heightToCap = float64(highest) - human.base
		default:
			fmt.Printf("human.Init: Name: \"%s\"\n", name)
		}
		if err != nil {
			return
		}
		human.Node.Add(mesh)
	}
	return nil
}

var NewMeshSkin = func(
	dec       *obj.Decoder,
	skin      *texture.Texture2D,
	skinColor *math32.Color4,
	underwear *texture.Texture2D,
	uwFabric  *math32.Color4,
	uwDetail  *math32.Color4,
	uwTrim    *math32.Color4,
	object    *obj.Object,
) (*graphic.Mesh, error) {
	geom, err := dec.NewGeometry(object)
	if err != nil {
		return nil, err
	}
	mat := new(HumanSkinMaterial)
	mat.Init("HumanSkin", &math32.Color4{1, 0, 1, 1})
	mat.Udata.SkinColor = skinColor
	mat.Udata.UwFabric = uwFabric
	mat.Udata.UwDetail = uwDetail
	mat.Udata.UwTrim = uwTrim
	mat.AddTexture(skin)
	mat.AddTexture(underwear)
	mesh := graphic.NewMesh(geom, mat)
	return mesh, nil
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
