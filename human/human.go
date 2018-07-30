package human

import (
	"fmt"
	"strings"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/loader/obj"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
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

func New(dec *obj.Decoder) (human *Human, err error) {
	human = new(Human)
	err = human.Init(dec)
	return
}

func (human *Human) Base() float64 {
	return human.base
}

func (human *Human) HeightToCap() float64 {
	return human.heightToCap
}

func (human *Human) HeightToEye() float64 {
	return human.heightToEye
}

func (human *Human) Init(dec *obj.Decoder) (err error) {
	human.Node = core.NewNode()
	for idx := 0; idx < len(dec.Objects); idx++ {
		var mesh *graphic.Mesh
		name := dec.Objects[idx].Name
		switch {
		case strings.HasSuffix(name, "-highpolyeyes"):
			mesh, err = dec.NewMesh(&dec.Objects[idx])
			human.Eyes = mesh
			lowest, highest, ok := ofsRange(dec, &dec.Objects[idx], 2)
			if ! ok {
				fmt.Printf("No vertices in \"%s\" (1)\n", name)
			}
			human.heightToEye = (float64(lowest) + float64(highest)) * 0.5
			frontest, _, ok2 := ofsRange(dec, &dec.Objects[idx], 1)
			if ! ok2 {
				fmt.Printf("No vertices in \"%s\" (2)\n", name)
			}
			human.frontOfEye = float64(frontest)
		case strings.HasSuffix(name, "-female_generic"):
			fallthrough
		case strings.HasSuffix(name, "-male_generic"):
			mesh, err = NewMeshSkin(dec, &dec.Objects[idx])
			human.Skin = mesh
			lowest, highest, ok := ofsRange(dec, &dec.Objects[idx], 2)
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

var NewMeshSkin = func(dec *obj.Decoder, object *obj.Object) (*graphic.Mesh, error) {
	geom, err := dec.NewGeometry(object)
	if err != nil {
		return nil, err
	}
	mat := new(material.Standard)
	mat.Init("HumanSkin", math32.NewColor("magenta"))
	mat.AddTexture(NakedSkin)
	mat.AddTexture(Underwear)
	mesh := graphic.NewMesh(geom, mat)
	return mesh, nil
}

func ofsRange(dec *obj.Decoder, object *obj.Object, ofs int) (lowest, highest float32, ok bool) {
	for _, face := range object.Faces {
		for _, vertex := range face.Vertices {
			val := dec.Vertices[vertex + ofs]
			if ok {
				if val < lowest {
					lowest = val
				}
				if val > highest {
					highest = val
				}
			} else {
				lowest = val
				highest = val
				ok = true
			}
		}
	}
	return
}
