package state

import (
	"github.com/amyadzuki/amygolib/str"

	"github.com/g3n/engine/window"
)

type State struct {
	fnCurrent, fnNext func(*State)
	fns               map[string]func(*State)
	state             string
	win               *window.Window
}

func New(win *window.Window) *State {
	s := new(State)
	s.Init(win)
	return s
}

func (s *State) Init(win *window.Window) {
	s.SetWindow(win)
}

func (s *State) OnEnter(name string, cb func(*State)) {
	s.fns[str.Simp(name) + "{"] = cb
}

func (s *State) OnLeave(name string, cb func(*State)) {
	s.fns[str.Simp(name) + "}"] = cb
}

func (s *State) Register(name string, cb func(*State)) {
	s.fns[str.Simp(name)] = cb
}

func (s *State) Run() {
	s.fnCurrent = s.fnNext
	enter, eok := s.fns[s.state + "{"]
	leave, lok := s.fns[s.state + "}"]
	if eok {
		enter(s)
	}
	for !s.win.ShouldClose() && s.fnNext == s.fnCurrent {
		s.current(s)
	}
	if lok {
		leave(s)
	}
}

func (s *State) SetNext(state string) {
	if fn, ok := s.fns[state], ok {
		s.state = state
		s.fnNext = fn
	} else {
		panic("Unregistered state: \"" + state + "\"")
	}
}

func (s *State) SetWindow(win *window.Window) {
	s.win = win
}
