package main

import (
	"fmt"
	"sync"
	"time"
)

func waitGroups() {
	const noWorkers = 5
	var waitGroup sync.WaitGroup

	for i := 0; i < noWorkers; i++ {
		waitGroup.Add(1) // adding counter when new goroutine is started.

		go func() {
			defer waitGroup.Done() // decreasing counter when goroutine is finished
			startWorker(i)
		}()
	}

	waitGroup.Wait()
	fmt.Printf("main thread synchronised, waited for %d workers finishing", noWorkers)
}

func startWorker(i int) {
	fmt.Printf("starting worker %d\n", i)
	time.Sleep(time.Second)
	fmt.Printf("terminating worker %d\n", i)
}
