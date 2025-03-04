package main

import "fmt"

func channels() {
	var messages = make(chan string)

	go func() { messages <- "ping" }()

	var ping = <-messages

	fmt.Println(ping)
}
