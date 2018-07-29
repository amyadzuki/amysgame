package events

var Event map[Id]Mapping
var EventName map[Id]string

func init() {
	Event = make(map[Id]Mapping)
	EventName = make(map[Id]string)
	EventName[ToggleFullScreen] = "Toggle FullScreen"
	EventName[ToggleInventory] = "Inventory"
}

type Id int
const (
	nothing Id = iota
	ToggleFullScreen
	ToggleInventory
)
