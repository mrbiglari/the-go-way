package main

import "fmt"

func channels() {
	messages := make(chan string) // an unbuffered channel (make(chan T)) requires both the sender and receiver to be ready at the same time; the sender blocks until the receiver reads the value, ensuring strict synchronization.

	go func() {
		messages <- "ping" /* blocks here until a receiver reads from the channel */
	}()

	ping := <-messages // is blocked until a sender sends a value into the channel

	fmt.Println(ping)
}
