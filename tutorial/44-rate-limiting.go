package main

import (
	"fmt"
	"time"
)

func rateLimiting() {
	unbufferedRateLimiting()
	bufferedRateLimiting()
}

func bufferedRateLimiting() {
	const burstBuffer = 3
	burstLimiter := make(chan time.Time, burstBuffer)
	for i := 0; i < burstBuffer; i++ {
		burstLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(500 * time.Millisecond) {
			burstLimiter <- t
		}
	}()

	burstRequests := generateRequests(5)
	for request := range burstRequests {
		<-burstLimiter
		fmt.Println("request: ", request, time.Now())
	}
}

func unbufferedRateLimiting() {
	const totalRequests = 5
	requests := generateRequests(totalRequests)
	limiter := time.Tick(500 * time.Millisecond)
	for request := range requests {
		<-limiter
		fmt.Println("request:", request, time.Now())
	}
}

func generateRequests(requestsCount int) chan int {
	requests := make(chan int, requestsCount)
	for i := 0; i < requestsCount; i++ {
		requests <- i
	}
	close(requests)
	return requests
}
