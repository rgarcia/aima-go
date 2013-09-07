package minimax

import (
	"fmt"
	"github.com/rgarcia/aima-go/adversarial"
	"log"
	"math"
)

var Debug = false

func debug(s string) {
	if Debug {
		log.Printf("minimax: %s\n", s)
	}
}

// Given the current state of the game, returns the optimal minimax action
func Decide(state adversarial.State) adversarial.Action {
	action, _ := MaxValue(state)
	return action
}

func MaxValue(state adversarial.State) (adversarial.Action, float64) {
	if state.Terminal() {
		return nil, state.Utility()
	}
	var max *adversarial.Successor
	var maxvalue = math.Inf(-1)
	successors := state.Successors()
	for i, successor := range successors {
		_, value := MinValue(successor.State)
		if max == nil || maxvalue < value {
			max = &successors[i]
			maxvalue = value
		}
	}
	if max == nil {
		panic(fmt.Sprintf("Non-terminal state has zero successors %v", state))
	}
	debug(fmt.Sprintf("MaxValue %v => (%v, %f)", state, max.Action, maxvalue))
	return max.Action, maxvalue
}

func MinValue(state adversarial.State) (adversarial.Action, float64) {
	if state.Terminal() {
		return nil, state.Utility()
	}
	var min *adversarial.Successor
	var minvalue = math.Inf(1)
	successors := state.Successors()
	for i, successor := range successors {
		_, value := MaxValue(successor.State)
		if min == nil || minvalue > value {
			min = &successors[i] // can't do &successor #wtfgo
			minvalue = value
		}
	}
	if min == nil {
		panic(fmt.Sprintf("Non-terminal state has zero successors %v", state))
	}
	debug(fmt.Sprintf("MinValue %v => (%v, %f", state, min.Action, minvalue))
	return min.Action, minvalue
}
