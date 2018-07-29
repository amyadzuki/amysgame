package events

type Mapping struct {
	Name string
	Func func(string, string, interface{})
}

func (m *Mapping) Defer(evname string, event interface{}) {
	fn := m.Func
	if fn != nil {
		fn(m.Name, evname, event)
	}
}
