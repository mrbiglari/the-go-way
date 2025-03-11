package main

import (
	"fmt"
	"maps"
)

func _maps() {
	bob := "Bob"
	sydney := "Sydney"

	scores := make(map[string]float32)
	scores[bob] = 6.9
	scores[sydney] = 4.2

	fmt.Printf("%s got a score of %.1f in the test.\n", bob, scores[bob])
	fmt.Printf("%s got a score of %.1f in the test.\n", sydney, scores[sydney])

	_, hasKey := scores[bob]
	if hasKey {
		fmt.Printf("The map has a score for bob.\n")
	}

	delete(scores, bob)
	if _, hasKey := scores[bob]; hasKey {
		fmt.Printf("Bob's score has been removed.\n")
	}

	clear(scores)
	if len(scores) == 0 {
		fmt.Printf("The map has been cleared.\n")
	}

	aliases := map[string]string{"apollo13": "bob", "artemis": "sydney"}
	aliasesDuplicated := map[string]string{"apollo13": "bob", "artemis": "sydney"}
	if maps.Equal(aliases, aliasesDuplicated) {
		fmt.Printf("The following two maps are equal: %v, %v", aliases, aliasesDuplicated)
	}
}
