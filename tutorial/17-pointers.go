package main

import "fmt"

func setZeroValueViaValue(variable int) {
	variable = 0
}

func setZeroValueViaPointer(variable *int) {
	*variable = 0
}

func pointers() {
	var x = 5                                 // type of `x` is `int`
	fmt.Printf("Value of 'x' is: %d\n", x)    // 3
	fmt.Printf("Address of 'x' is: %d\n", &x) // 0x42131100

	var y = &x                              // & operator gets `address of` `x` // type of `y` is `*int`
	fmt.Printf("Value of 'y' is: %d\n", y)  // 0x42131100 // the value is the address of `x`.
	fmt.Printf("Value of 'x' is: %d\n", *y) // 3 // * operator `dereferences` pointer `y` to get value of `x`.

	*y = 10 // dereferencing `y` allows to update value of `x`
	fmt.Printf("Value of 'x' is now: %d\n", x)
	fmt.Printf("Value of 'x' is now: %d\n", *y)

	var z *int // z is a pointer to an int (not initialized yet, nil)
	if z == nil {
		fmt.Println("Uninitialised pointer to int `z` is `nil`.")
	}

	var variable = 72

	setZeroValueViaValue(variable)
	fmt.Println("Value of `variable` is: ", variable) // 72

	setZeroValueViaPointer(&variable)
	fmt.Println("Value of `variable` is: ", variable) // 0
}
