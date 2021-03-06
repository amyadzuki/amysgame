package events

import (
	"github.com/g3n/engine/window"
)

var Keycode map[window.Key]Id

func init() {
	Keycode = make(map[window.Key]Id)
	Keycode[window.KeyF12] = ToggleCharaDesigner
	Keycode[window.KeyI] = ToggleInventory
	Keycode[window.KeyF11] = ToggleFullScreen
}

func OnKeyboardKey(evname string, event interface{}) {
	ev, ok := event.(*window.KeyEvent)
	if !ok {
		return
	}
	var down bool
	switch ev.Action {
	case window.Press:
		down = true
	case window.Release:
		down = false
	default:
		return
	}
	id, ok := Keycode[ev.Keycode]
	if !ok {
		return
	}
	m, ok := Event[id]
	if !ok {
		return
	}
	m.Defer(id, down, evname, event)
}
