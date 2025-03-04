package main

import "fmt"

func ping(_ping chan<- string, message string) { // ping is a receive-only channel
	_ping <- message
}

func pong(_ping <-chan string, _pong chan<- string) { // pong is a send-only channel
	var message = <-_ping
	_pong <- message
}

func channelDirections() {
	var message = "hello world!"
	var pingChannel = make(chan string, 1)
	ping(pingChannel, message)

	var pongChannel = make(chan string, 1)
	pong(pingChannel, pongChannel)

	fmt.Println(<-pongChannel)
}
