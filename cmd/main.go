package main

import (
	"fmt"
	lollipop "lollipop/internal"
)

type TripStatus string

const (
	TripStatusScheduled  TripStatus = "SCHEDULED"
	TripStatusInProgress TripStatus = "IN_PROGRESS"
	TripStatusCompleted  TripStatus = "COMPLETED"
	TripStatusCancelled  TripStatus = "CANCELLED"
)

func main() {
	stateMachine := lollipop.NewStateMachine(TripStatusScheduled)

	stateMachine.AddTransition(TripStatusScheduled, TripStatusInProgress)
	stateMachine.AddTransition(TripStatusScheduled, TripStatusCancelled)
	stateMachine.AddTransition(TripStatusInProgress, TripStatusCompleted)

	fmt.Printf("Initial state: %v\n", stateMachine.State)

	if err := stateMachine.Transition(TripStatusInProgress); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("State after transition: %s\n", stateMachine.State)

	if err := stateMachine.Transition(TripStatusCancelled); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
