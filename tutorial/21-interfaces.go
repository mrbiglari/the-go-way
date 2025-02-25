package main

import (
	"fmt"
	"math"
)

// interface example 1
type shape interface {
	area() float64
	perimeter() float64
}

// square
type square struct {
	side float64
}

func (v square) area() float64 {
	return v.side * v.side
}
func (v square) perimeter() float64 {
	return 4 * v.side
}

// circle
type circle struct {
	radius float64
}

func (v circle) area() float64 {
	return math.Pi * v.radius * v.radius
}
func (v circle) perimeter() float64 {
	return 2 * math.Pi * v.radius
}

// interface example 2
type animal interface {
	makeSound()
}

type cat struct{}

func (p cat) makeSound() { // if a method has a value receiver (T), both a value (T) and a pointer (&T) implement the interface.
	fmt.Println("Meow!")
}

type dog struct{}

func (p *dog) makeSound() { // if a method has a pointer receiver (*T), only a pointer (&T) implements the interface.
	fmt.Println("Woof Woof!")
}

/*
When to Use Which?

✔ Use a value receiver (T) if:
    ● The method does not modify the struct.
    ● The struct is small, and copying it is cheap.
    ● You want both values (T) and pointers (&T) to implement the interface.

✔ Use a pointer receiver (*T) if:
    ● The method modifies the struct.
    ● The struct is large, and passing a pointer avoids unnecessary copying.
    ● You want only pointers (&T) to implement the interface.
*/

func interfaces() {
	var shapes = []shape{square{side: 10}, circle{radius: 3.0}}
	for _, shape := range shapes {
		measure(shape)
	}

	var animals = []animal{cat{}, &dog{}} // cat struct as value, dog struct as pointer
	for _, animal := range animals {
		animal.makeSound()
	}
}

func measure(shape shape) {
	fmt.Println(shape)
	fmt.Println(shape.area())
	fmt.Println(shape.perimeter())
}
