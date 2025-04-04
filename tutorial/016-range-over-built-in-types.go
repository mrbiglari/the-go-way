package main

import (
	"fmt"
	"slices"
)

func rangeOverBuiltInTypes() {
	numbers := []int{2, 3, 4}
	for index, number := range numbers {
		fmt.Printf("number[%d]:%d\n", index, number)
	}

	chunkMe := []string{"It", "is", "possible", "to", "range", "over", "chunks", "of", "a", "slice", "using", "Chunk()"}
	for chunk := range slices.Chunk(chunkMe, 3) {
		fmt.Println(chunk)
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
