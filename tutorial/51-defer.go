package main

import "fmt"

/*
✔ Notes on the 'defer' mechanism:
	● deferred calls execute after function finishes but before the function exits
	● it's useful for cleanup operations (like closing files, releasing locks, etc.)
	● defer statements are executed in LIFO (Last-In, First-Out) order

✔ Output:
	1. start function...
	2. end function...
	3. post-function: pre-clean up task 1...
	4. post-function: pre-clean up task 2...
	6. start clean up...
	7. finish clean up...
	8. post clean up...
*/

func cleanUp() {
	fmt.Println("6. start clean up...")
	defer fmt.Println("8. post clean up...")
	fmt.Println("7. finish clean up...")
}

func _defer() {
	fmt.Println("1. start function...")

	defer cleanUp()
	defer fmt.Println("4. post-function: pre-clean up task 2...")
	defer fmt.Println("3. post-function: pre-clean up task 1...")

	fmt.Println("2. end function...")
}
