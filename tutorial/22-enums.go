package main

import "fmt"

type ServerState int

const (
	Idle ServerState = iota
	Connected
	Error
)

var stateToTextMapping = map[ServerState]string{
	Idle:      "idle",
	Connected: "connected",
	Error:     "error",
}

func (v ServerState) String() string {
	return stateToTextMapping[v]
}

func enums() {
	serverState := Idle

	if serverState == Idle {
		print(serverState)
	}

	nextState := transition(serverState)
	print(nextState)
}

func print(state ServerState) {
	fmt.Printf("The server's status is '%v' (value: %d).\n", state, state)
}

func transition(state ServerState) ServerState {
	switch state {
	case Idle:
		return Connected
	case Connected:
		return Idle
	case Error:
		return Error
	default:
		panic(fmt.Errorf("unknown state: %s", state))
	}
}
