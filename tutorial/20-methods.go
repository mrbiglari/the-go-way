package main

import "fmt"

type rectangle struct {
	width, height int
}

// method with a pointer receiver
func (self *rectangle) area() int {
	return self.width * self.height
}

// method with a value receiver
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
	rectangleValue := rectangle{width: 5, height: 3}
	fmt.Println("area:", rectangleValue.area())
	fmt.Println("perimeter:", rectangleValue.perimeter())

	rectanglePointer := &rectangleValue
	fmt.Println("area:", rectanglePointer.area())
	fmt.Println("perimeter:", rectanglePointer.perimeter())

	rectangleValue.updateValueState(10, 15)
	fmt.Println("rectangle dimensions are still:", rectangleValue)

	rectanglePointer.updatePointerState(10, 15)
	fmt.Println("rectangle dimensions are now:", rectangleValue)
	fmt.Println("rectangle dimensions are now:", rectanglePointer)
}
