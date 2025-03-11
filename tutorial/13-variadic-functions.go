package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for num := range nums {
		total += num
	}
	return total
}

func variadicFunctions() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum := sum(nums...)

	fmt.Println("Total sum is:", sum)
}
