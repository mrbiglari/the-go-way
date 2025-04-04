package main

import (
	"fmt"
	"time"
)

func do(channel chan<- int, factor int) {
	time.Sleep(time.Duration(factor) * time.Second)
	channel <- factor
}

func _select() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go do(channel1, 1)
	go do(channel2, 2)

	for i := 0; i < 2; i++ { // since there are two goroutines sending messages, we need to loop twice to receive both.
		select { // select waits until at least one channel is ready.
		case message := <-channel1:
			fmt.Println("received", message) // once a case is executed, 'select' exits and does not wait for the other channels, to process multiple messages, 'select' must be inside a loop.
		case message := <-channel2:
			fmt.Println("received", message)
		}
	}
}
