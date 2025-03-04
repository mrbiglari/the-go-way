package main

import "fmt"

func channelBuffering() {
	var channel = make(chan string, 2)
	channel <- "first"
	channel <- "second"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
