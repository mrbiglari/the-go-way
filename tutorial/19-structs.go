package main

import "fmt"

//The capitalisation of an identifier (type, method, function, variable, or struct field) determines its visibility within and outside a package.
type person struct {
	name string
	age  int
}

// Constructor returning a pointer
func newPerson(name string) *person { // returns a pointer
	var person = person{name: name}
	person.age = 42
	return &person // return address of
}

// Constructor returning a value
func makePerson(name string) person { // returns a value
	var person = person{name: name}
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
	var bob = person{name: "Bob"}
	fmt.Println(bob)

	var sydney = person{name: "Sydney", age: 45}
	fmt.Println(sydney)
	fmt.Printf("%s is %d years old.", sydney.name, sydney.age)

	var trevor = &person{name: "Trevor", age: 22}
	fmt.Println(trevor)

	var amanda = newPerson("Amanda")
	fmt.Println(amanda)

	var buddy = makePerson("Buddy")
	fmt.Println(buddy)

	var dog = struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
