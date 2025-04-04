package main

import "fmt"

func rangeOverChannels() {
	channel := make(chan int, 3)

	channel <- 1
	channel <- 2
	channel <- 3

	close(channel)

	for integer := range channel {
		fmt.Println(integer)
	}
}
