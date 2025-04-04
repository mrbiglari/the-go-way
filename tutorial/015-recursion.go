package main

import "fmt"

func fibonacci(n int) int {
	if n < 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func recursion() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
}
