package main

import (
	"errors"
	"fmt"
)

func decrement(argument int) (int, error) {
	if argument == 69 {
		return -1, errors.New("this number doesn't need decrementing") // built-in Go standard library function for creating string-based errors.
	}
	return argument - 1, nil
}

var errOutOfTea = fmt.Errorf("no more tea available")
var errOutOfPower = fmt.Errorf("water isn't boiling")

func makeTea(argument int) error {
	if argument == 2 {
		return errOutOfTea
	} else if argument == 4 {
		return fmt.Errorf("making tea... %w", errOutOfPower)
	}
	return nil
}

func _errors() {
	var numbers = []int{13, 69}
	for _, number := range numbers {
		if value, _error := decrement(number); _error != nil {
			fmt.Printf("Failed to decrement value of %d.\n", number)
		} else {
			fmt.Printf("Successfully decremented value %d to %d.\n", number, value)
		}
	}

	for i := range 6 {
		var _error = makeTea(i)
		if _error != nil {
			if errors.Is(_error, errOutOfTea) {
				fmt.Printf("%v, we should buy new tea!\n", _error)
			} else if errors.Is(_error, errOutOfPower) {
				fmt.Printf("%v... oops no electricity!\n", _error)
			} else {
				fmt.Println("Unknown error.")
			}
		} else {
			fmt.Println("No error occurred.")
		}
	}
}
