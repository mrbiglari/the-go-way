package main

import "fmt"

func channelBuffering() {
	var channel = make(chan string, 2) // a buffered channel (make(chan T, N)) allows the sender to send up to N values without blocking, only blocking if the buffer is full, and the receiver only blocks when the buffer is empty.
	channel <- "first"
	channel <- "second"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
