package alphabeta

import (
	"github.com/rgarcia/aima-go/examples"
	"testing"
)

func TestSixFour(t *testing.T) {
	Debug = true
	initialState := examples.FigSixFourState{}
	action := Decide(initialState)
	if action.Label() != "a1" {
		t.Errorf("Expected a1, got %s", action.Label())
	}
}
