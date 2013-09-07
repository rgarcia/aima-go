package alphabeta

import (
	"fmt"
	"github.com/rgarcia/aima-go/adversarial"
	"log"
	"math"
)

var Debug = false

func debug(s string) {
	if Debug {
		log.Printf("alphabeta: %s\n", s)
	}
}

// Given the current state of the game, returns the optimal minimax action
func Decide(state adversarial.State) adversarial.Action {
	action, _ := MaxValue(state, math.Inf(-1), math.Inf(+1))
	return action
}

func MaxValue(state adversarial.State, alpha float64, beta float64) (adversarial.Action, float64) {
	if state.Terminal() {
		return nil, state.Utility()
	}
	var max *adversarial.Successor
	var maxvalue = math.Inf(-1)
	successors := state.Successors()
	for i, successor := range successors {
		_, value := MinValue(successor.State, alpha, beta)
		if max == nil || maxvalue < value {
			max = &successors[i]
			maxvalue = value
		}
		if maxvalue >= beta {
			debug(fmt.Sprintf("MaxValue %v pruning at %v => (%v, %f)", state, successor, max.Action, maxvalue))
			return max.Action, maxvalue
		}
		alpha = math.Max(alpha, maxvalue)
	}
	if max == nil {
		panic(fmt.Sprintf("Non-terminal state has zero successors %v", state))
	}
	debug(fmt.Sprintf("MaxValue %v => (%v, %f)", state, max.Action, maxvalue))
	return max.Action, maxvalue
}

func MinValue(state adversarial.State, alpha float64, beta float64) (adversarial.Action, float64) {
	if state.Terminal() {
		return nil, state.Utility()
	}
	var min *adversarial.Successor
	var minvalue = math.Inf(1)
	successors := state.Successors()
	for i, successor := range successors {
		_, value := MaxValue(successor.State, alpha, beta)
		if min == nil || minvalue > value {
			min = &successors[i] // can't do &successor #wtfgo
			minvalue = value
		}
		if minvalue <= alpha {
			debug(fmt.Sprintf("MinValue %v pruning at %v => (%v, %f)", state, successor, min.Action, minvalue))
			return min.Action, minvalue
		}
		beta = math.Min(beta, minvalue)
	}
	if min == nil {
		panic(fmt.Sprintf("Non-terminal state has zero successors %v", state))
	}
	debug(fmt.Sprintf("MinValue %v => (%v, %f", state, min.Action, minvalue))
	return min.Action, minvalue
}
