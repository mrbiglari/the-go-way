package main

import (
	"fmt"
	"time"
)

func _switch() {
	i := 2
	switch i {
	case 1:
		fmt.Println("i is one")
		// break not needed as it breaks by default.
	case 2:
		fmt.Println("i is two")
		fallthrough // use fallthrough if you do want multiple cases to execute.
	case 3:
		fmt.Println("i is either two or three")
	}

	day := "Saturday"
	switch day {
	case "Saturday", "Sunday": // acts like (day == "Saturday" || day == "Sunday")
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	time := time.Now()
	switch { // switch without an expression behaves as switch true, allowing logical expressions.
	case time.Hour() < 12 || time.Hour() == 12: // cases can have conditional expressions like if statements
		fmt.Println("It's before noon or noon.")
	default:
		fmt.Println("It's after noon.")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
