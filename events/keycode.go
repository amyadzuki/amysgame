package events

import (
	"github.com/g3n/engine/window"
)

var Keycode map[int]EventId

func init() {
	Keycode = make(map[int]EventId)
	Keycode[window.KeyI] = Inventory
	Keycode[window.KeyF11] = ToggleFullscreen
}

func OnKeyboardKey(evname string, event interface{}) {
	ev, ok := event.(*window.KeyEvent)
	if !ok {
		return
	}
	kei, ok := Keycode[ev.Keycode]
	if !ok {
		return
	}
	m, ok := KevEvent[kei]
	if !ok {
		return
	}
	m.Defer(evname, event)
}
