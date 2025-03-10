package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func atomicCounters() {
	var counter atomic.Uint64
	var wait sync.WaitGroup

	for i := 0; i < 50; i++ {
		wait.Add(1)
		go incrementThousandTimes(&counter, &wait)
	}

	wait.Wait()

	fmt.Println("counter value: ", counter.Load())
}

func incrementThousandTimes(counter *atomic.Uint64, wait *sync.WaitGroup) {
	defer wait.Done()
	for i := 0; i < 1000; i++ {
		counter.Add(1)
	}
}
