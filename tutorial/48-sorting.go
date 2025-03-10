package main

import (
	"fmt"
	"slices"
)

func sorting() {
	strings := []string{"c", "b", "a"}
	slices.Sort(strings)
	fmt.Println(strings)

	numbers := []int{7, 3, 1}
	slices.Sort(numbers)
	fmt.Println(numbers)

	fmt.Printf("Is slices:%v sorted? %v", numbers, slices.IsSorted(numbers))
}
