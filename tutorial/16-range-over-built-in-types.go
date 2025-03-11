package main

import "fmt"

func rangeOverBuiltInTypes() {
	numbers := []int{2, 3, 4}
	for index, number := range numbers {
		fmt.Printf("number[%d]:%d\n", index, number)
	}

	aliases := map[string]string{"apollo13": "bob", "artemis": "sydney"}
	for key, value := range aliases {
		fmt.Printf("%s -> %s\n", key, value)
	}

	for key := range aliases {
		fmt.Println("key:", key)
	}

	for index, character := range "go" {
		fmt.Println(index, character)
	}
}
