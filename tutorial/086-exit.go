package main

import (
	"fmt"
	"os"
)

func _exit() {
	fmt.Println("I'm gonna print!")
	defer fmt.Println("No you're not!")
	os.Exit(3) // code zero indicates success, non-zero an error
}
