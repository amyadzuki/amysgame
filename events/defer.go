package events

func defer(name, evname string, event interface{}, fn func(string, string, interface{})) {
	if fn != nil {
		fn(name, evname, event)
		return
	}
}
