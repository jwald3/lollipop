package main

import (
	"fmt"
	lollipop "lollipop/internal"
)

func main() {
	stateMachine := lollipop.NewStateMachine("A")

	stateMachine.AddTransition("A", "B")
	stateMachine.AddTransition("B", "C")
	stateMachine.AddTransition("C", "A")

	fmt.Printf("Initial state: %s\n", stateMachine.State)

	if err := stateMachine.Transition("B"); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("State after transition: %s\n", stateMachine.State)

	if err := stateMachine.Transition("A"); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
