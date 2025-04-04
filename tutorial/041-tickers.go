package main

import (
	"fmt"
	"time"
)

func tick(ticker *time.Ticker, done chan bool) {
	for {
		select {
		case <-done:
			return // without this, the tick goroutine keeps waiting forever, even after stopping the ticker.
		case time := <-ticker.C:
			fmt.Println("ticked at: ", time)
		}
	}
}

func stopTickerAfterDuration(seconds int, ticker *time.Ticker, done chan bool) {
	timer := time.NewTimer(time.Duration(seconds) * time.Second)
	<-timer.C
	ticker.Stop()
	done <- true // the done channel allows a signal to be sent, telling the goroutine to stop gracefully.
	fmt.Println("ticking stopped")
}

func tickers() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go tick(ticker, done)
	stopTickerAfterDuration(3, ticker, done)
}
