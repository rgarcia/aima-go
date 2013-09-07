package adversarial

type Action interface {
	Label() string
}

type Successor struct {
	Action
	State
}

// States are nodes in the search tree.
type State interface {
	Successors() []Successor
	Terminal() bool
	Utility() float64
}
