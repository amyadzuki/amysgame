package events

import (
	"github.com/g3n/engine/window"
)

var Button map[int]EventId

func init() {
	Button = make(map[int]EventId)
}

func OnMouseButton(evname string, event interface{}) {
	ev, ok := event.(*window.MouseEvent)
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
	id, ok := Button[ev.Button]
	if !ok {
		return
	}
	m, ok := KevEvent[id]
	if !ok {
		return
	}
	m.Defer(down, evname, event)
}
