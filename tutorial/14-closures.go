package main

import "fmt"

func getCounterIncrementFunc() func() int {
	var counter = 0

	var incrementFunc = func() int {
		counter++
		return counter
	}
	return incrementFunc
}

func closures() {
	var incrementFunc = getCounterIncrementFunc()

	fmt.Println("Initial counter value is:", 0)

	for i := 0; i < 5; i++ {
		fmt.Println("...incrementing, counter value is now:", incrementFunc())
	}

	fmt.Println("\nLet's start over!\n")

	var anotherIncrementFunc = getCounterIncrementFunc() // Each closure in Go retains access to variables from its outer scope, encapsulating them within its execution context.
	fmt.Println("Initial counter value is:", 0)

	for i := 0; i < 3; i++ {
		fmt.Println("...incrementing, counter value is now:", anotherIncrementFunc())
	}
}
