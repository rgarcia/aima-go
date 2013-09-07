package examples

import "github.com/rgarcia/aima-go/adversarial"

// Fig 6.4
type FigSixFourAction struct {
	Action string
}

func (a FigSixFourAction) Label() string {
	return a.Action
}

type FigSixFourState struct {
	depth  int
	action FigSixFourAction
	value  float64
}

func (state FigSixFourState) String() string {
	switch state.depth {
	case 0:
		return "A"
	case 1:
		switch state.action.Label() {
		case "a1":
			return "B"
		case "a2":
			return "C"
		case "a3":
			return "D"
		}
	}
	return "unknown"
}

func (state FigSixFourState) Terminal() bool {
	return state.depth == 2
}

func (state FigSixFourState) Utility() float64 {
	return state.value
}

func (state FigSixFourState) Successors() []adversarial.Successor {
	switch state.depth {

	case 0:
		return []adversarial.Successor{
			{FigSixFourAction{"a1"}, FigSixFourState{state.depth + 1, FigSixFourAction{"a1"}, 0}},
			{FigSixFourAction{"a2"}, FigSixFourState{state.depth + 1, FigSixFourAction{"a2"}, 0}},
			{FigSixFourAction{"a3"}, FigSixFourState{state.depth + 1, FigSixFourAction{"a3"}, 0}},
		}

	case 1:
		switch state.action.Label() {
		case "a1":
			return []adversarial.Successor{
				{FigSixFourAction{"b1"}, FigSixFourState{state.depth + 1, FigSixFourAction{"b1"}, 3}},
				{FigSixFourAction{"b2"}, FigSixFourState{state.depth + 1, FigSixFourAction{"b2"}, 12}},
				{FigSixFourAction{"b3"}, FigSixFourState{state.depth + 1, FigSixFourAction{"b3"}, 8}},
			}
		case "a2":
			return []adversarial.Successor{
				{FigSixFourAction{"c1"}, FigSixFourState{state.depth + 1, FigSixFourAction{"c1"}, 2}},
				{FigSixFourAction{"c2"}, FigSixFourState{state.depth + 1, FigSixFourAction{"c2"}, 4}},
				{FigSixFourAction{"c3"}, FigSixFourState{state.depth + 1, FigSixFourAction{"c3"}, 6}},
			}
		case "a3":
			return []adversarial.Successor{
				{FigSixFourAction{"d1"}, FigSixFourState{state.depth + 1, FigSixFourAction{"d1"}, 14}},
				{FigSixFourAction{"d2"}, FigSixFourState{state.depth + 1, FigSixFourAction{"d2"}, 5}},
				{FigSixFourAction{"d3"}, FigSixFourState{state.depth + 1, FigSixFourAction{"d3"}, 2}},
			}
		}
	}
	panic("unreachable")
	return []adversarial.Successor{}
}
