package main

import "fmt"

//The capitalisation of an identifier (type, method, function, variable, or struct field) determines its visibility within and outside a package.
type person struct {
	name string
	age  int
}

// Constructor returning a pointer
func newPerson(name string) *person { // returns a pointer
	person := person{name: name}
	person.age = 42
	return &person // return address of
}

// Constructor returning a value
func makePerson(name string) person { // returns a value
	person := person{name: name}
	person.age = 42
	return person // return value
}

/*
When to Use Which?

✔ Use a pointer (*person) when:
    ● The struct is large, and you want to avoid copying it.
    ● You need to modify the original struct after returning it — stateful struct instances.
    ● You want to share the same instance across multiple functions.

✔ Use a value (person) when:
    ● The struct is small, and copying it is cheap.
    ● You do not need to modify the original after returning — stateless struct instances.
    ● You want immutable data (every function gets a copy).
*/

func structs() {
	bob := person{name: "Bob"}
	fmt.Println(bob)

	sydney := person{name: "Sydney", age: 45}
	fmt.Println(sydney)
	fmt.Printf("%s is %d years old.", sydney.name, sydney.age)

	trevor := &person{name: "Trevor", age: 22}
	fmt.Println(trevor)

	rob := new(person) // Similar to `&person{}` The `new(person)` function allocates memory for a variable of a given type and returns a pointer to it. While this type of initialisation is possible, most idiomatic Go code prefers composite literals over the use of `new`.
	rob.name = "Rob"
	fmt.Println(rob)

	amanda := newPerson("Amanda")
	fmt.Println(amanda)

	buddy := makePerson("Buddy")
	fmt.Println(buddy)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
