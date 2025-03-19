package main

import (
	"flag"
	"fmt"
)

func commandlineFlags() {
	print := fmt.Println

	word := flag.String("word", "bob", "a string")
	number := flag.Int("num", 69, "an integer")
	fork := flag.Bool("fork", false, "a boolean")
	var svar string
	flag.StringVar(&svar, "svar", "sydney", "a string variable")

	flag.Parse()

	print(*word)
	print(*number)
	print(*fork)
	print(svar)
	print(flag.Args())
}
