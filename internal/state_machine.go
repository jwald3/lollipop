package lollipop

import "fmt"

type State any

type StateMachine struct {
	State       State
	Transitions map[State][]State
}

func NewStateMachine(initialState State) *StateMachine {
	return &StateMachine{
		State:       initialState,
		Transitions: make(map[State][]State),
	}
}

// add transitions to the state machine's registry. if a state is not present in the map of
// transitions, we will add it and its "to" state
func (sm *StateMachine) AddTransition(from, to State) {
	if sm.Transitions[from] == nil {
		sm.Transitions[from] = []State{}
	}
	sm.Transitions[from] = append(sm.Transitions[from], to)
}

func (sm *StateMachine) CanTransition(to State) bool {
	validStates, exists := sm.Transitions[sm.State]
	// if the current state isn't included in the transaction definitions, you cannot
	// transition to any state.
	if !exists {
		return false
	}

	// loop over the valid transition options until a match or the end of the list
	for _, validState := range validStates {
		if validState == to {
			return true
		}
	}

	return false
}

func (sm *StateMachine) Transition(to State) error {
	if !sm.CanTransition(to) {
		return fmt.Errorf("invalid transition from %v to %v", sm.State, to)
	}

	sm.State = to
	return nil
}
