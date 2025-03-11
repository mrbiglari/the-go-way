package main

import "fmt"

func ifElse() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	if 8%2 == 0 && 4%2 == 0 {
		fmt.Println("both 8 and 4 are even")
	}

	b := 10
	if i := b; b == 10 {
		fmt.Println(i, "is 10")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
