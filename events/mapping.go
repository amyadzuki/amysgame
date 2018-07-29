package events

type Mapping struct {
	Name string
	Func func(string, bool, string, interface{})
}

func (m *Mapping) Defer(down bool, evname string, event interface{}) {
	fn := m.Func
	if fn != nil {
		fn(m.Name, down, evname, event)
	}
}
