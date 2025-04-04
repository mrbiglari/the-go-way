package main

import "fmt"

func panics() {
	panic("a problem")
}

func recovers() {
	if _error := recover(); _error != nil { // the return value of recover is the error raised in the call to panic
		fmt.Printf("recovered from '%v'", _error)
	}
}

func _recover() {
	defer recovers() // recover must be called within a deferred function. When the enclosing function panics, the defer will activate and a recover call within it will catch the panic
	panics()

	fmt.Println("I'm not used to being ignored.") // this line is never executed
}
