package main

import (
	"fmt"
	"reflect"
)

func reflection() {
	typeOf()

	structValueFields()
	structPointerFields()

	getStructValueMethods()
	getStructPointerMethods()

	structValueConstructors()
	structPointerConstructors()

	structValueMethodsInvoke()
	structPointerMethodsInvoke()
}

func typeOf() {
	var x int = 42
	var y float64 = 3.14
	var z string = "Hello"

	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(y))
	fmt.Println(reflect.TypeOf(z))

	// Get only the name of the type
	fmt.Println(reflect.TypeOf(x).Name())
	fmt.Println(reflect.TypeOf(y).Name())
	fmt.Println(reflect.TypeOf(z).Name())
}

type candy struct {
	brand string
}

func structValueFields() {
	candy := candy{brand: "Mars"}
	structType := reflect.TypeOf(candy)
	structValue := reflect.ValueOf(candy)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		value := structValue.Field(i)
		fmt.Printf("Field: %s, Value: %s, Type: %s\n", field.Name, value, field.Type)
	}
}

func structPointerFields() {
	candy := &candy{brand: "Bounty"}
	structType := reflect.TypeOf(candy).Elem()
	structValue := reflect.ValueOf(candy).Elem()
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		value := structValue.Field(i)
		fmt.Printf("Field: %s, Value: %s, Type: %s\n", field.Name, value, field.Type)
	}
}

type bird struct{}

func (self bird) Sing() { // value receiver method
}

func (self *bird) Fly() { // pointer receiver method
}

func getStructValueMethods() {
	value := reflect.TypeOf(bird{}) // struct as a value
	for i := 0; i < value.NumMethod(); i++ {
		fmt.Println(value.Method(i).Name) // getting methods from a value instance: Sing
	}
}

func getStructPointerMethods() {
	pointer := reflect.TypeOf(&bird{}) // struct as a pointer
	for i := 0; i < pointer.NumMethod(); i++ {
		fmt.Println(pointer.Method(i).Name) // getting methods from a pointer instance: Sing & Fly
	}
}

type payment struct {
	id     string
	amount int
}

func makePayment(id string, amount int) payment {
	return payment{id: id, amount: amount}
}

func newPayment(id string, amount int) *payment {
	return &payment{id: id, amount: amount}
}

func structValueConstructors() {
	constructor := reflect.ValueOf(makePayment)                            // get function by reflection
	args := []reflect.Value{reflect.ValueOf("12345"), reflect.ValueOf(30)} // call the function with arguments
	payment := constructor.Call(args)[0].Interface().(payment)
	fmt.Println(payment)
}

func structPointerConstructors() {
	constructor := reflect.ValueOf(newPayment)                             // get function by reflection
	args := []reflect.Value{reflect.ValueOf("12345"), reflect.ValueOf(30)} // call the function with arguments
	payment := constructor.Call(args)[0].Interface().(*payment)
	fmt.Println(payment)
}

type squirrel struct {
	hunger int
}

func newSquirrel() *squirrel {
	squirrel := squirrel{}
	squirrel.hunger = 20
	return &squirrel
}

type nuts = int

func (self *squirrel) Eat(amount nuts) { // MethodByName Only Works on Exported Methods, Methods must start with an uppercase letter to be accessible via reflection.
	self.hunger -= amount
}

func (self squirrel) Munch(amount nuts) { // value receiver method (works on copies)
	self.hunger -= amount
}

func structValueMethodsInvoke() {
	squirrel := newSquirrel()
	fmt.Println("Starting Hunger Levels:", squirrel.hunger) // 20

	valueType := reflect.TypeOf(*squirrel) // get type of `squirrel` (not `*squirrel`)

	method, exists := valueType.MethodByName("Munch")
	if !exists {
		fmt.Println("Method not found!")
		return
	}

	args := []reflect.Value{reflect.ValueOf(*squirrel), reflect.ValueOf(5)} // first argument must be a value (not a pointer)

	method.Func.Call(args) // call the method dynamically

	fmt.Println("Final Hunger Levels:", squirrel.hunger) // still 20 because `Eat` modified a copy
}

func structPointerMethodsInvoke() {
	squirrel := newSquirrel()
	fmt.Println("Starting Hunger Levels:", squirrel.hunger) // 20

	pointerType := reflect.TypeOf(squirrel)

	method, exists := pointerType.MethodByName("Eat")

	if !exists {
		fmt.Println("No method called 'Eat' found on the type 'squirrel'.")
		return
	}

	args := []reflect.Value{reflect.ValueOf(squirrel), reflect.ValueOf(5)} // first argument must be the receiver (squirrel)

	method.Func.Call(args)                               // call the method dynamically
	fmt.Println("Final Hunger Levels:", squirrel.hunger) // 15
}
