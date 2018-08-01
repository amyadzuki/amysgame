package state

type IState interface {
	Common() *State
	Enter()
	Leave()
	Main()
}

func (s IState) Run() {
	common := s.Common()
	common.current = common.next
	s.Enter()
	for !common.win.ShouldClose() && common.next == common.current {
		s.Main()
	}
	s.Leave()
}

func (s IState) SetNext(next int) {
	s.Common().next = next
}

func (s IState) SetWindow(win *window.Window) {
	s.Common().win = win
}

type State struct {
	win     *window.Window
	current int32
	next    int32
}

func (s *State) Common() *State {
	return s
}

func (s *State) Enter() {
}

func (s *State) Init(win *window.Window, state int) {
	s.win = win
	s.current = int32(state)
	s.next = int32(state)
}

func (s *State) Leave() {
}

func (s *State) Main() {
}
