package main

import (
	"fmt"
	lollipop "lollipop/internal"
)

type LightState string

const (
	Off LightState = "OFF"
	On  LightState = "ON"
)

func main() {
	sm := lollipop.NewStateMachine(Off)

	// register valid state paths
	sm.AddTransition(Off, On)
	sm.AddTransition(On, Off)

	// register entry actions for the available states
	sm.SetEntryAction(On, func() error {
		fmt.Println("Light bulb warming up...")
		return nil
	})

	sm.SetEntryAction(Off, func() error {
		fmt.Println("Light bulb cooling down...")
		return nil
	})

	fmt.Printf("Current state: %v\n", sm.State)

	// perform transactions with entry actions attached
	fmt.Println("Turning light on...")
	// if the transition is valid, the state machine will call the entry action
	// associated with entering the "On" state
	if err := sm.Transition(On); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Current state: %v\n", sm.State)

	fmt.Println("Turning light off...")
	if err := sm.Transition(Off); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Current state: %v\n", sm.State)
}
