package main

import "fmt"

func send(jobs chan<- int) {
	for i := 0; i < 5; i++ {
		jobs <- i
		fmt.Printf("job %d sent\n", i)
	}
	close(jobs) // closing the channel allows 'receive' to break from its loop
	fmt.Println("sent all jobs")
}

func receive(jobs <-chan int, done chan bool) {
	for {
		job, hasMore := <-jobs
		if hasMore {
			fmt.Println("received job", job)
		} else { // loop breaks when jobs channel is closed
			fmt.Println("received all jobs")
			done <- true
			break
		}
	}
}

func closingChannels() {
	jobs := make(chan int)
	done := make(chan bool)

	go send(jobs)
	go receive(jobs, done)

	<-done              // blocks until done receives a value
	_, isOpen := <-jobs // reading from a closed channel succeeds immediately, returning the zero value of the underlying type. The optional second return value is true if the value received was delivered by a successful send operation to the channel, or false if it was a zero value generated because the channel is closed and empty.
	if !isOpen {
		fmt.Println("can't receive more jobs. Is the channel still open?", isOpen)
	}
}
