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

// State represents any value that can be used as a state
type State any

// Action is a function that is executed when entering or exiting a state
type Action func() error

// StateMachine manages state transitions and their associated actions
type StateMachine struct {
	State        State
	Transitions  map[State][]State
	InitialState State
	EntryActions map[State]Action
	ExitActions  map[State]Action
}

func NewStateMachine(initialState State) *StateMachine {
	return &StateMachine{
		State:        initialState,
		Transitions:  make(map[State][]State),
		InitialState: initialState,
		EntryActions: make(map[State]Action),
		ExitActions:  make(map[State]Action),
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

func (sm *StateMachine) SetEntryAction(state State, action Action) {
	sm.EntryActions[state] = action
}

func (sm *StateMachine) SetExitAction(state State, action Action) {
	sm.ExitActions[state] = action
}

func (sm *StateMachine) Reset() {
	sm.State = sm.InitialState
}
