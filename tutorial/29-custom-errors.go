package main

import (
	"errors"
	"fmt"
)

type customError struct {
	argument int
	message  string
}

func (pointer *customError) Error() string { // A custom error type can be created by implementing the Error() method on them.
	return fmt.Sprintf("%d - %s", pointer.argument, pointer.message)
}

func increment(argument int) (int, error) {
	if argument == 69 {
		return -1, &customError{argument: argument, message: "This number doesn't need incrementing!"}
	}
	return argument + 1, nil
}

func customErrors() {
	var _, _error = increment(69)
	var _customError *customError

	if errors.As(_error, &_customError) {
		fmt.Printf("%d, %s", _customError.argument, _customError.message)
	} else {
		fmt.Println("Provided error is not of customError type.")
	}
}
