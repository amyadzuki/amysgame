package events

func defer(name, evname string, event interface{}, fn, fallback func(string, string, interface{})) {
	if fn != nil {
		fn(name, evname, event)
		return
	}
	if fallback != nil {
		fn(name, evname, event)
		return
	}
}
