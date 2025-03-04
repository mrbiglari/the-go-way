package main

import (
	"fmt"
	"time"
)

func someLogic(done chan bool) {
	fmt.Print("working...")
	time.Sleep(2 * time.Second)

	done <- true
}

func channelSynchronization() {
	var done = make(chan bool)
	go someLogic(done)

	<-done
	fmt.Println("done")
}
