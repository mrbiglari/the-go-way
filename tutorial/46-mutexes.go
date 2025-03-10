package main

import (
	"fmt"
	"sync"
)

type keyedCounters struct {
	mutex       sync.Mutex
	countersMap map[string]int
}

func (self *keyedCounters) increment(key string) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	self.countersMap[key]++
}

func mutexes() {
	counters := &keyedCounters{countersMap: map[string]int{"a": 0, "b": 0}}
	var wait sync.WaitGroup

	wait.Add(3)
	go updateContainer("a", 100000, counters, &wait)
	go updateContainer("a", 100000, counters, &wait)
	go updateContainer("b", 100000, counters, &wait)
	wait.Wait()
	fmt.Println(counters.countersMap)
}

func updateContainer(key string, n int, counters *keyedCounters, wait *sync.WaitGroup) {
	defer wait.Done()
	for i := 0; i < n; i++ {
		counters.increment(key)
	}
}
