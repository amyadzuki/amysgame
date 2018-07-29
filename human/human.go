package human

import (
	"fmt"
	"strings"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/loader/obj"
	"github.com/g3n/engine/material"
)

type Human struct {
	Node *core.Node

	Eyes *graphic.Mesh
	Skin *graphic.Mesh

	heightToEye float64
}

func New(dec *obj.Decoder) (human *Human, err error) {
	human = new(Human)
	err = human.Init(dec)
	return
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
			human.Eyes = &dec.Objects[idx]
			mesh, err = dec.NewMesh(human.Eyes)
			index := dec.Objects[idx].Faces[0].Vertices[0]
			_ = index // TODO
			// human.heightToEye = TODO:
		case strings.HasSuffix(name, "-female_generic"):
			fallthrough
		case strings.HasSuffix(name, "-male_generic"):
			human.Skin = &dec.Objects[idx]
			mesh, err = NewMeshSkin(dec, human.Skin)
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
