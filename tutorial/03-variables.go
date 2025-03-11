package main

import "fmt"

func variables() {
	var a = "initial"
	fmt.Println(a)

	var b, c = 1, 2
	fmt.Println(b, c)

	var d int = 4
	fmt.Println(d)

	var e, f int = 5, 6
	fmt.Println(e, f)

	var g = true
	fmt.Println(g)

	var h bool
	fmt.Println(h)

	i := "short"
	fmt.Println(i)
}
