package statemachine

import (
	"errors"
	"fmt"
)

// Common errors that may be returned by the state machine
var (
	ErrInvalidTransition = errors.New("invalid state transition")
	ErrEntryActionFailed = errors.New("entry action failed")
	ErrExitActionFailed  = errors.New("exit action failed")
)

// State represents any value that can be used as a state - you are expected to enforce a valid
// set of valid state options in your implementation
type State any

// Action is a function that is executed when entering or exiting a state
type Action func() error

// StateMachine manages state transitions and their associated actions
type StateMachine struct {
	State        State             // a reference to the current state at a given time
	Transitions  map[State][]State // each state and the valid, allowed transitions
	InitialState State             // the state used in `Reset()` calls
	EntryActions map[State]Action  // the functions called when entering a state
	ExitActions  map[State]Action  // the functions called when exiting a state
}

func NewStateMachine(initialState State) *StateMachine {
	return &StateMachine{
		State:        initialState,
		Transitions:  make(map[State][]State), // These properties use methods to set their values explicitly.
		InitialState: initialState,
		EntryActions: make(map[State]Action), // ---
		ExitActions:  make(map[State]Action), // ---
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

// go from one state to another, performing exit and entry actions where applicable.
// the transition only sets the state machine's current status, so any intention to
// use a state machine to update an object's status requires the use of entry/exit actions
func (sm *StateMachine) Transition(to State) error {
	if !sm.CanTransition(to) {
		return fmt.Errorf("%w: from %v to %v", ErrInvalidTransition, sm.State, to)
	}

	if exitAction := sm.ExitActions[sm.State]; exitAction != nil {
		if err := exitAction(); err != nil {
			return fmt.Errorf("%w: %v", ErrExitActionFailed, err)
		}
	}

	oldState := sm.State
	sm.State = to

	if entryAction := sm.EntryActions[to]; entryAction != nil {
		if err := entryAction(); err != nil {
			sm.State = oldState
			return fmt.Errorf("%w: %v", ErrEntryActionFailed, err)
		}
	}

	return nil
}

// Set or replace the entry action for a given state. The entry action is a generic function that
// you will define in your implementation. This is called during the transition following the state machine
// transitioning from the present to the destination state
func (sm *StateMachine) SetEntryAction(state State, action Action) {
	sm.EntryActions[state] = action
}

// Set or replace the exit action for a given state. The exit action is a generic function that
// you will define in your implementation. This is called during the transition prior to the state machine
// transitioning from the present to the destination state
func (sm *StateMachine) SetExitAction(state State, action Action) {
	sm.ExitActions[state] = action
}

func (sm *StateMachine) Reset() {
	sm.State = sm.InitialState
}
