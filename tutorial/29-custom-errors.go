package main

import (
	"errors"
	"fmt"
)

type customError struct {
	argument int
	message  string
}

func (self *customError) Error() string { // error is an interface, therefore a custom error type can be created by implementing the Error() method on a struct.
	return fmt.Sprintf("%d - %s", self.argument, self.message)
}

func increment(argument int) (int, error) {
	if argument == 69 {
		return -1, &customError{argument: argument, message: "This number doesn't need incrementing!"}
	}
	return argument + 1, nil
}

func customErrors() {
	_, _error := increment(69)
	var _customError *customError

	if errors.As(_error, &_customError) {
		fmt.Printf("%d, %s", _customError.argument, _customError.message)
	} else {
		fmt.Println("Provided error is not of customError type.")
	}
}
