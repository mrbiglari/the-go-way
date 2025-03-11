package main

import (
	"fmt"
	"time"
)

func count() {
	for i := 0; i < 5; i++ {
		fmt.Printf("...%d\n", i)
	}
}

func goRoutines() {
	fmt.Println("Let's start counting!")

	go count()

	interject := func() {
		fmt.Println("No! I interject!")
	}
	go interject()

	time.Sleep(time.Second)
	fmt.Println("Counting finished!")
}
