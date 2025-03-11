package main

import "fmt"

func getThree() (int, string) {
	return 3, "three"
}

func multipleReturnValues() {
	threeInt, threeString := getThree()
	fmt.Printf("Return values, '%d' and '%s' fetched from getThree() function invocation.", threeInt, threeString)
}
