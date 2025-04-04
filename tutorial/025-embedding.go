package main

import "fmt"

type base struct {
	number int
}

func (self base) describe() string {
	return fmt.Sprintf("base with num:%v", self.number)
}

type container struct {
	base // Go's embedding is somewhere between composition (has-a) and inheritance—it reuses the fields, but without creating an "is-a" relationship" // it looks like inheritance because of field promotion and it works like composition because it doesn't allow upcasting.
	text string
}

type describer interface {
	describe() string
}

func embedding() {
	container := container{base: base{number: 20}, text: "Bob"}
	fmt.Println(container.base.number, container.number, container.text)
	fmt.Println("describe:", container.describe())

	var describer describer = container
	fmt.Println("describe:", describer.describe())
}
