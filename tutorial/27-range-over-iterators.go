package main

import (
	"fmt"
	"iter"
)

type ChainList[T any] struct {
	head *Chain[T]
}

type Chain[T any] struct {
	value T
	next  *Chain[T]
}

func (pointer *ChainList[T]) enumerate() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := pointer.head; i != nil; i = i.next {
			if !yield(i.value) {
				return
			}
		}
	}
}

func fibonacciSequence() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1
		for {
			if !yield(a) { // yield is an internal `iter` function that determines if the iteration should continue, true: continues and false: terminates.
				return
			}
			a, b = b, a+b
		}
	}
}

func iterators() {
	list := &ChainList[int]{head: &Chain[int]{value: 1, next: &Chain[int]{value: 2, next: &Chain[int]{value: 3}}}}

	for element := range list.enumerate() {
		fmt.Println(element)
	}

	for n := range fibonacciSequence() {
		if n > 15 {
			break // when for is terminated (break or return), yield() returns false and terminates generating iteration items.
		}
		fmt.Println(n)
	}
}
