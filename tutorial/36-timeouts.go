package main

import (
	"fmt"
	"time"
)

func waitSeconds(channel chan bool, factor int) {
	time.Sleep(time.Duration(factor) * time.Second)
	channel <- true
}

func timeouts() {
	channel1 := make(chan bool)
	go waitSeconds(channel1, 2) // sends value to channel after 2 seconds.

	select {
	case <-channel1:
		fmt.Println("message received after two seconds.")
	case <-time.After(time.Second): // times out after one second before value is sent.
		fmt.Println("timed out before receiving the message.")
	}

	channel2 := make(chan bool)
	go waitSeconds(channel2, 1)

	select {
	case <-channel2: // message is received after one second.
		fmt.Println("message received after one second.")
	case <-time.After(3 * time.Second):
		fmt.Println("timed out before receiving the message.")
	}
}
