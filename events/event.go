package events

var Event map[Id]Mapping

func init() {
	Event = make(map[Id]Mapping)
	Event[ToggleFullScreen] = Mapping{"Toggle FullScreen", nil}
	Event[ToggleInventory] = Mapping{"Inventory", nil}
}

type Id int
const (
	nothing Id = iota
	ToggleFullScreen
	ToggleInventory
)
