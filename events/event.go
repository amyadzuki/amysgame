package events

var Event map[Id]Mapping

func init() {
	Event = make(map[Id]Mapping)
	Event[ToggleFullScreen] = Mapping{"Toggle FullScreen", &OnToggleFullScreen}
	Event[ToggleInventory] = Mapping{"Inventory", &OnToggleInventory}
}

type Id int
const (
	nothing Id = iota
	ToggleFullScreen
	ToggleInventory
)

var (
	OnToggleFullScreen &func(string, bool, string, interface{})
	OnToggleInventory &func(string, bool, string, interface{})
)
