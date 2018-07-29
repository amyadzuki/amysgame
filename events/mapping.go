package events

type Mapping func(int, bool, string, interface{})

func (m Mapping) Defer(id Id, down bool, evname string, event interface{}) {
	if m != nil {
		m(int(id), down, evname, event)
	}
}
