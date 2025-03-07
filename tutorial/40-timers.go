package main

import (
	"fmt"
	"time"
)

func sprint(timer *time.Timer) {
	<-timer.C
	fmt.Println("Timer fired")
}

func timers() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go sprint(timer2)
	hasStopped := timer2.Stop() // it returns true if the call stops the timer, false if the timer has already
	// expired or been stopped.
	if hasStopped {
		fmt.Println("Timer 2 has stopped")
	}
}
