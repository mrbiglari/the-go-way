package main

import (
	"fmt"
	"slices"
)

func _slices() {
	array := [3]int{1, 2, 3}
	slicedArray := array[:] // create a slice from an array
	fmt.Println(slicedArray)

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	words := []string{"Bob", "went", "to", "Sydney"} // defining a slice with dynamic size.
	fmt.Println(words, "len:", len(words), "cap:", cap(words))

	words = append(words, "last", "year") // note that we need to accept a return value from append as we may get a new slice value.
	fmt.Println(words, "len:", len(words), "cap:", cap(words))

	moreWords := make([]string, 4, 10) // defining a slice with set length and capacity for better optimisation (optional).
	moreWords[0] = "Sydney"
	moreWords[1] = "went"
	moreWords[2] = "to"
	moreWords[3] = "Bob Creek"
	fmt.Println(moreWords, "len:", len(moreWords), "cap:", cap(moreWords))

	repeatWords := make([]string, len(moreWords), cap(moreWords))
	copy(repeatWords, moreWords)
	repeatWords = append(repeatWords, "again!")
	fmt.Println(repeatWords, "len:", len(repeatWords), "cap:", cap(repeatWords))

	counting := []string{"zero", "one", "two", "three", "four", "five"}

	twoTillFour := counting[2:5]
	fmt.Printf("Let's count from two till four: %v\n", twoTillFour)

	fromThreeOnward := counting[3:]
	fmt.Printf("Let's count from three onward: %v\n", fromThreeOnward)

	tillThree := counting[:4]
	fmt.Printf("Let's count till three: %v\n", tillThree)

	B1 := []string{"Are", "you", "thinking", "what", "I'm", "thinking", "B1?"}
	B2 := []string{"Are", "you", "thinking", "what", "I'm", "thinking", "B1?"}

	if slices.Equal(B1, B2) {
		fmt.Println(B1)
		fmt.Println("I think I am B2!")
	}

	length := 3
	twoDimensionalArray := make([][]int, length)
	for i := 0; i < length; i++ {
		twoDimensionalArray[i] = make([]int, length)
		for j := 0; j < length; j++ {
			twoDimensionalArray[i][j] = 0 // Set values to zero.
		}
	}
	fmt.Println("Two-dimensional array holding zero values:", twoDimensionalArray)
}
