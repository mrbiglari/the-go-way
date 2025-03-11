package main

import "fmt"

type number interface { // type constraint
	int | int64 | float32 | float64
}

func add[T number](a, b T) T {
	return a + b
}

type LinkedList[T any] struct {
	head, tail *Node[T] // before generics was added to Go, interface{} would be substituted however this meant each node could get a different value type linkedList.Push("text"), linkedList.Push(24). This is not ideal as we may want to know (or enforce) types at compile time and not run time.
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

func (self *LinkedList[T]) Push(value T) {
	if self.tail == nil {
		self.head = &Node[T]{value: value}
		self.tail = self.head
	} else {
		self.tail.next = &Node[T]{value: value}
		self.tail = self.tail.next
	}
}

func (self *LinkedList[T]) GetAll() []T {
	var nodes []T
	for node := self.head; node != nil; node = node.next {
		nodes = append(nodes, node.value)
	}
	return nodes
}

type stringer interface {
	string() string
}

type logger interface {
	log()
}

type loggable interface {
	stringer
	logger
}

type item struct {
	name string
}

func (self item) string() string {
	return "item: " + self.name
}

func (self item) log() {
	fmt.Println("Logging:", self.string())
}

func printAndLog[T loggable](item T) { // functionally no different than func printAndLog(item loggable) {...} // using a generic function (T Loggable) makes sense only if you need to: (1) know the exact (concrete) type of item at compile time and (2) if you need to return T or use it in a way that depends on its concrete type.
	fmt.Println("Printing:", item.string())
	item.log()
}

func generics() {
	fmt.Println(add(3, 5))
	fmt.Println(add(3.4, 5.6))

	linkedList := &LinkedList[int]{}
	linkedList.Push(5)
	linkedList.Push(6)
	linkedList.Push(7)
	fmt.Println(linkedList.GetAll())

	dice := item{name: "dice"}
	printAndLog(dice)
}
