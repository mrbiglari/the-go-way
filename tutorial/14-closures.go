package main

import "fmt"

func getCounterIncrementFunc() func() int {
	counter := 0

	incrementFunc := func() int {
		counter++
		return counter
	}
	return incrementFunc
}

func closures() {
	incrementFunc := getCounterIncrementFunc()

	fmt.Println("Initial counter value is:", 0)

	for i := 0; i < 5; i++ {
		fmt.Println("...incrementing, counter value is now:", incrementFunc())
	}

	fmt.Println("Let's start over!")

	anotherIncrementFunc := getCounterIncrementFunc() // Each closure in Go retains access to variables from its outer scope, encapsulating them within its execution context.
	fmt.Println("Initial counter value is:", 0)

	for i := 0; i < 3; i++ {
		fmt.Println("...incrementing, counter value is now:", anotherIncrementFunc())
	}
}
