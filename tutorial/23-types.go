package main

import "fmt"

type movie struct { // 1. user-defined struct type
	name string
}

type age int // 2. user-defined type

type intAlias = int // 3. type alias

type speaker interface { // 4. user-defined interface type
	Speak() string
}

type function func(int, int) int // 5. user-defined function type

func types() {
	movie := movie{name: "The Matrix"}
	fmt.Println(movie)

	var age age = 25 // not interchangeable with the original type int
	fmt.Println(age)

	var integer intAlias = 34 // interchangeable with the original type int
	fmt.Println(integer)

	var speaker speaker
	if speaker == nil {
		fmt.Println("speaker is nil!")
	}

	var add function = func(a, b int) int { return a + b }
	sum := add(5, 3)
	fmt.Println(sum)
}
