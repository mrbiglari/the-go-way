package main

import "fmt"

type A interface{ A() string }
type B interface{ B() string }
type C interface {
	A
	B
	C() string
}
type D struct{ s string }

func (d *D) A() string { return "A" }
func (d *D) B() string { return "B" }
func (d *D) C() string { return "C" }

func variableAssignments() {
	var a11 int // declaration without assignment
	a11 = 42
	var a12 int = 42 // delectation explicit // most idiomatic outside functions
	var a13 = 42     // declaration inferred
	a14 := 42        // short-form, same as `var a12 = 42` // most idiomatic inside functions

	var a21 float64 = 42 // when the `float64` is needed instead of the default `int`
	a22 := 42.0          // a26 is a float64 due to fractional value being specified

	bob := &D{"bob"}
	var a31 C = bob
	var a32 A = bob
	var a33 B = bob
	var a34 A = a31
	var a35 B = a31

	var a41 *D
	a41 = &D{"bob"}

	var a42 D
	a42 = D{"bob"}

	var a43 *D = nil // same as `var a43 *D`
	a43 = &D{"bob"}

	var a44 *D = &D{"bob"}
	var a45 D = D{"bob"}

	var a46 = &D{"bob"}
	var a47 = D{"bob"}

	var a51 interface{}
	var a52 interface{} = nil
	var a53 interface{} = 53
	var a54 interface{} = "a54"

	fmt.Println(a11, a12, a13, a14,
		a21, a22,
		a31.C(), a32.A(), a33.B(), a34.A(), a35.B(),
		a41, a42, a43, a44, a45, a46, a47,
		a51, a52, a53, a54)
}
