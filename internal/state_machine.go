package lollipop

import "fmt"

type State any
type Action func() error

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
		return fmt.Errorf("invalid transition from %v to %v", sm.State, to)
	}

	// if the transaction is valid, check if there is a valid exit action function
	// if there is a exit action available, attempt to call it, returning an error if unsuccessful
	// if there's not an exit action, just skip
	if exitAction := sm.ExitActions[sm.State]; exitAction != nil {
		if err := exitAction(); err != nil {
			return fmt.Errorf("exit action failed: %w", err)
		}
	}

	// retain the old state in case we need to rollback
	oldState := sm.State

	sm.State = to

	// attempt to call the entry action if one is registered for the destination state.
	// if there is one, run the action, rolling back if unable to successfully call it
	if entryAction := sm.EntryActions[to]; entryAction != nil {
		if err := entryAction(); err != nil {
			sm.State = oldState
			return fmt.Errorf("entry action failed: %w", err)
		}
	}

	// at this point, the app called the exit action, performed the transition,
	// and called the entry action successfully.
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
