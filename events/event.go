package events

var Event map[EventId]Mapping

func init() {
	Event = make(map[EventId]Mapping)
	Event[ToggleFullScreen] = Mapping{"Toggle FullScreen", OnToggleFullScreen}
	Event[ToggleInventory] = Mapping{"Inventory", OnToggleInventory}
}

type EventId int
const (
	nothing EventId = iota
	ToggleFullScreen
	ToggleInventory
)

var (
	onNothing func(string, string, event interface{})
	OnToggleFullScreen
	OnInventory
)
