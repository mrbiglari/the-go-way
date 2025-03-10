package main

import "os"

func divide(x, y float64) float64 {
	if y == 0 {
		panic("cannot divide by zero")
	}
	return x / y
}

func _panic() {
	panic("a problem occurred!") // output: 'panic: a problem occurred!'

	divide(8, 0) // output: 'panic: cannot divide by zero'

	_, _error := os.Open("nonexistent.txt") // in Go, functions (whether custom-defined by developers or from the standard library) generally do not panic by default. instead, they follow the convention of returning an error value, which allows the caller to handle errors gracefully.
	if _error != nil {
		panic(_error) // output: 'panic: open nonexistent.txt: The system cannot find the file specified.'
	}
}
