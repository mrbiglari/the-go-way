package main

import "fmt"

type rectangle struct {
	width, height int
}

// method with a pointer receiver, for when mutable, shared-state behaviour is needed
func (self *rectangle) area() int {
	return self.width * self.height
}

// method with a value receiver, for when copy-only, read-only behaviour is needed
func (self rectangle) perimeter() int {
	return 2*self.width + 2*self.height
}

func (self *rectangle) updatePointerState(width, height int) {
	self.width = width // updates internal state
	self.height = height
}
func (self rectangle) updateValueState(width, height int) {
	self.width = width // updates local copy
	self.height = height
}

func methods() {
	value := rectangle{width: 5, height: 3}

	fmt.Println("area:", value.area())           // (&value).area() — Go takes address
	fmt.Println("perimeter:", value.perimeter()) // value.perimeter() — direct call

	pointer := &value

	fmt.Println("area:", pointer.area())           // pointer.area() — direct call
	fmt.Println("perimeter:", pointer.perimeter()) // (*pointer).perimeter() — Go dereferences

	value.updateValueState(10, 15) // value.updateValueState(...) — direct call on a copy
	fmt.Println("rectangle dimensions are still:", value)

	pointer.updatePointerState(10, 15) // pointer.updatePointerState(...) — direct call
	fmt.Println("rectangle dimensions are now:", value)
	fmt.Println("rectangle dimensions are now:", pointer)
}
