package events

var Event map[EventId]Mapping

func init() {
	Event = make(map[EventId]Mapping)
	Event[Inventory] = Mapping{"Inventory", OnInventory}
	Event[ToggleFullscreen] = Mapping{"Toggle Fullscreen", OnToggleFullscreen}
}

type EventId int
const (
	nothing EventId = iota
	Inventory
	ToggleFullscreen
)

var (
	onNothing func(string, string, event interface{})
	OnInventory
	OnToggleFullscreen
)
