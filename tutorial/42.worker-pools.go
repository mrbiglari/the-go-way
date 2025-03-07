package main

import (
	"fmt"
	"time"
)

func work(jobs <-chan int, results chan<- int, workerId int) {
	for job := range jobs {
		fmt.Printf("worker %d started job:%d\n", workerId, job)
		time.Sleep(time.Second)
		fmt.Printf("worker %d finished job:%d\n", workerId, job)
		results <- job
	}
}

func startWorkers(jobs <-chan int, results chan<- int) {
	for i := 0; i < 3; i++ { // start workers
		go work(jobs, results, i)
	}
}

func createJobs(jobs chan<- int, numJobs int) {
	for i := 0; i < numJobs; i++ {
		jobs <- i
	}
	close(jobs)
}

func workers() {
	const numJobs = 5
	jobs := make(chan int)
	results := make(chan int)

	startWorkers(jobs, results)

	go createJobs(jobs, numJobs)

	for i := 1; i <= numJobs; i++ {
		<-results // finally we collect all the results of the work. This ensures that the worker goroutines have finished otherwise they remain blocked.
		fmt.Println("terminating worker")
	}
}
