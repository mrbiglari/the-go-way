package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func lineFilters() {
	print := fmt.Println

	file, _ := os.Open("01-hello-world.go")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		upperCaseLine := strings.ToUpper(line)
		print(upperCaseLine)
	}

	if _error := scanner.Err(); _error != nil { // .Err() stores the first encountered error and stops scanning.
		print(_error)
		os.Exit(1)
	}
}
