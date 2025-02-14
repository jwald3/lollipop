![lollipop github](https://github.com/user-attachments/assets/08f458d3-0f6e-4a2b-8212-917f76e78d10)

# Lollipop

Lollipop is a lightweight, flexible state machine implementation in Go. It provides a simple way to manage state transitions in your applications with support for entry and exit actions.

## Features

- Simple and intuitive API
- Type-safe state transitions
- Entry and exit actions for states
- State transition validation
- Easy state machine reset capability
- Zero external dependencies

## Installation

```bash
go get github.com/jwald3/lollipop
```

## Usage

```go
import "github.com/jwald3/lollipop/pkg/statemachine"

func main() {
    // Create a new state machine
    sm := statemachine.NewStateMachine(InitialState)
    
    // Add transitions
    sm.AddTransition(InitialState, NextState)
    
    // Add actions
    sm.SetEntryAction(NextState, func() error {
        // Do something when entering NextState
        return nil
    })
}
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/yourusername/lollipop"
)

func main() {
    // Define your states (can be any type)
    const (
        Idle     = "IDLE"
        Running  = "RUNNING"
        Finished = "FINISHED"
    )

    // Create a new state machine with initial state
    sm := lollipop.NewStateMachine(Idle)

    // Define valid transitions
    sm.AddTransition(Idle, Running)
    sm.AddTransition(Running, Finished)
    sm.AddTransition(Finished, Idle)

    // Add entry and exit actions
    sm.SetEntryAction(Running, func() error {
        fmt.Println("Starting to run...")
        return nil
    })

    sm.SetExitAction(Running, func() error {
        fmt.Println("Finishing up...")
        return nil
    })

    // Perform transitions
    if err := sm.Transition(Running); err != nil {
        fmt.Printf("Error: %v\n", err)
    }

    if err := sm.Transition(Finished); err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}
```

## Detailed Usage

### Creating a State Machine

```go
// States can be of any type (string, int, custom type, etc.)
sm := lollipop.NewStateMachine(initialState)
```

### Defining Transitions

```go
// Add allowed transitions
sm.AddTransition(fromState, toState)

// Example with multiple transitions from one state
sm.AddTransition(StateA, StateB)
sm.AddTransition(StateA, StateC)
```

### Adding Actions

```go
// Add entry action for a state
sm.SetEntryAction(state, func() error {
    // Code to execute when entering this state
    return nil
})

// Add exit action for a state
sm.SetExitAction(state, func() error {
    // Code to execute when leaving this state
    return nil
})
```

### Performing Transitions

```go
// Check if transition is valid
if sm.CanTransition(newState) {
    // Perform the transition
    err := sm.Transition(newState)
    if err != nil {
        // Handle error
    }
}
```

### Resetting the State Machine

```go
// Reset to initial state
sm.Reset()
```

## Real-World Example: Document Processing

```go
package main

import (
    "fmt"
    "github.com/yourusername/lollipop"
)

type DocumentState string

const (
    Draft     DocumentState = "DRAFT"
    Review    DocumentState = "REVIEW"
    Approved  DocumentState = "APPROVED"
    Published DocumentState = "PUBLISHED"
)

func main() {
    // Create state machine for document workflow
    sm := lollipop.NewStateMachine(Draft)

    // Define valid transitions
    sm.AddTransition(Draft, Review)
    sm.AddTransition(Review, Draft)
    sm.AddTransition(Review, Approved)
    sm.AddTransition(Approved, Published)
    sm.AddTransition(Approved, Review)

    // Add entry actions
    sm.SetEntryAction(Review, func() error {
        fmt.Println("Notifying reviewers...")
        return nil
    })

    sm.SetEntryAction(Published, func() error {
        fmt.Println("Publishing document to website...")
        return nil
    })

    // Example workflow
    fmt.Printf("Current state: %v\n", sm.State)
    
    _ = sm.Transition(Review)
    fmt.Printf("Current state: %v\n", sm.State)
    
    _ = sm.Transition(Approved)
    fmt.Printf("Current state: %v\n", sm.State)
    
    _ = sm.Transition(Published)
    fmt.Printf("Current state: %v\n", sm.State)
}
```

## Error Handling

Lollipop provides comprehensive error handling for invalid transitions and failed actions:

```go
err := sm.Transition(newState)
switch {
case err != nil:
    if errors.Is(err, lollipop.ErrInvalidTransition) {
        // Handle invalid transition
    } else {
        // Handle other errors (like failed actions)
    }
default:
    // Transition successful
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

