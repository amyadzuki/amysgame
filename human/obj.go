package human

import (
	"fmt"
	"strings"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/loader/obj"
)

func NewGroup(dec *obj.Decoder) (*core.Node, error) {
	group := core.NewNode()
	for idx := 0; idx < len(dec.Objects); idx++ {
		mesh, err := dec.NewMesh(&dec.Objects[idx])
		if err != nil {
			return nil, err
		}
		name := dec.Objects[idx].Name
		fmt.Printf("(NewGroup: Name: \"%s\")\n", name)
		switch {
		case strings.HasPrefix(name, "Eye_"):
		case strings.HasPrefix(name, "young_"):
		default:
			fmt.Printf("NewGroup: Name: \"%s\"\n", name)
		}
		group.Add(mesh)
	}
	return group, nil
}
