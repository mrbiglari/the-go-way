package main

import "fmt"

/*
✔ Unbuffered channel (make(chan T)):
   ● It checks if a receiver is ready.
   ● If no receiver is ready, it blocks until one is.
   ● If a receiver is ready, it sends immediately without using any buffer.
✔ Buffered channel (make(chan T, N)):
   ● It checks if there is space in the buffer.
   ● If the buffer is not full, it places the value inside the buffer without blocking.
   ● If the buffer is full, it blocks until space is available.
*/

func channelBuffering() {
	var channel = make(chan string, 2) // a buffered channel (make(chan T, N)) allows the sender to send up to N values without blocking, only blocking if the buffer is full, and the receiver only blocks when the buffer is empty.
	channel <- "first"
	channel <- "second"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
