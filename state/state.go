package state

type State interface {
	Enter()
	Leave()
	Run()
}
