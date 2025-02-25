package main

import "fmt"

type rectangle struct {
	width, height int
}

// method with a pointer receiver
func (p *rectangle) area() int {
	return p.width * p.height
}

// method with a value receiver
func (v rectangle) perimeter() int {
	return 2*v.width + 2*v.height
}

func (p *rectangle) updatePointerState(width, height int) {
	p.width = width // updates internal state
	p.height = height
}
func (v rectangle) updateValueState(width, height int) {
	v.width = width // updates local copy
	v.height = height
}

func methods() {
	var rectangleValue = rectangle{width: 5, height: 3}
	fmt.Println("area:", rectangleValue.area())
	fmt.Println("perimeter:", rectangleValue.perimeter())

	var rectanglePointer = &rectangleValue
	fmt.Println("area:", rectanglePointer.area())
	fmt.Println("perimeter:", rectanglePointer.perimeter())

	rectangleValue.updateValueState(10, 15)
	fmt.Println("rectangle dimensions are still:", rectangleValue)

	rectanglePointer.updatePointerState(10, 15)
	fmt.Println("rectangle dimensions are now:", rectangleValue)
	fmt.Println("rectangle dimensions are now:", rectanglePointer)
}
