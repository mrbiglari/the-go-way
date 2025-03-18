package main

import (
	"fmt"
	"os"
)

/*
	$ go build main.go
	[./main a b c d]
	./main
	[a b c d]
	c
*/

func commandlineArguments() {
	print := fmt.Println

	print(os.Args)

	programmePath := os.Args[0]
	print(programmePath)

	arguments := os.Args[1:]
	print(arguments)

	print(os.Args[3])
}
