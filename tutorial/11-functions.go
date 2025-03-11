package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return plus(plus(a, b), c)
}

func functions() {
	noApples := 8
	noOranges := 7
	fmt.Printf("There are a total of %d Apples and Oranges.\n", plus(noApples, noOranges))

	noPears := 5
	fmt.Printf("There are a total of %d fruits.\n", plusPlus(noApples, noOranges, noPears))
}
