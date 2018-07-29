package human

import (
	"fmt"
	"strings"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/loader/obj"
)

type Human struct {
	Node *core.Node

	Eyes *obj.Object
	Skin *obj.Object

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

func (human *Human) Init(dec *obj.Decoder) error {
	human.Node = core.NewNode()
	for idx := 0; idx < len(dec.Objects); idx++ {
		mesh, err := dec.NewMesh(&dec.Objects[idx])
		if err != nil {
			return err
		}
		name := dec.Objects[idx].Name
		switch {
		case strings.HasSuffix(name, "-highpolyeyes"):
			human.Eyes = &dec.Objects[idx]
			index := dec.Objects[idx].Faces[0].Vertices[0]
			_ = index // TODO
			// human.heightToEye = TODO:
		case strings.HasSuffix(name, "-female_generic"):
			fallthrough
		case strings.HasSuffix(name, "-male_generic"):
			human.Skin = &dec.Objects[idx]
		default:
			fmt.Printf("human.Init: Name: \"%s\"\n", name)
		}
		human.Node.Add(mesh)
	}
	return nil
}
