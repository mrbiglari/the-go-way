package main

import (
	"fmt"
	"unsafe"
)

/*
	Type conversion uses the syntax: T(v), where T is the target type and v is the value to be converted.

	In Go, assignment and type conversion are distinct operations with strict type rules:

	1. Assignment (declaration form): `var x T = v`
	- Valid only if `v` is directly assignable to type `T`.
	- No type conversion occurs.

	2. Assignment (short form): `x := v`
	- Similar to the above, infers type from `v`.
	- Requires `v` to be assignable to the inferred type.

	3. Type conversion (declaration form): `var x T = T(v)`
	- Required when `v` is not assignable to `T`.
	- Performs an explicit type conversion at compile time.

	4. Type conversion (short form): `x := T(v)`
	- Shorthand for the above, also requires explicit conversion.
	- Used when `v` is not directly assignable to the inferred type.

	Note: Go does not support implicit type conversion. All conversions between non-assignable types must be explicit.
*/

var i int = 10

var a11 float64 = float64(i) // 1. basic conversion

type MyInt int    // user-defined type
type ThyInt = int // type alias

var a21 MyInt = MyInt(i) // 2. custom conversion
var a22 ThyInt = i       // assignment (not a conversion)

var a31 *int = &i // 3. pointer operation (not a conversion)
var a32 int = *a31

var a41 *float64 = (*float64)(unsafe.Pointer(&i)) // 4. unsafe conversion

var channel = make(chan int)
var a51 <-chan int = channel // 5. channel conversion // bidirectional to send-only
var a52 chan<- int = channel // bidirectional to receive-only

type Point struct{ x, y int }
type Coord struct{ x, y int }

var a61 Point = Point{x: 1, y: 2} // 6. Struct conversions
var a62 Coord = Coord(a61)
var a63 *Point = &Point{x: 1, y: 2}
var a64 *Coord = (*Coord)(a63)

type MyFunc func(int) int

func half(x int) int { return x / 2 }

var double func(int) int = func(x int) int { return x + x }
var a71 MyFunc = MyFunc(half) // 7. Function conversions
var a72 MyFunc = (MyFunc)(double)

type E interface{ E() string }
type F interface{ F() string }
type G interface {
	E
	F
	G() string
}
type H struct{ s string }

func (h *H) E() string { return "E" }
func (h *H) F() string { return "F" }
func (h *H) G() string { return "G" }

var h *H = &H{"bob"}
var a81 G = G(h) // 8. Interface conversions
var a82 E = E(h) // struct to interface
var a83 F = (F)(h)
var a84 E = E(a81) // interface to interface
var a85 F = F(a81)

func typeConversion() {
	// 1. Basic type conversions
	a12 := int(42)
	a13 := (int)(42) // brackets over type is optional
	a14 := float64(a12)
	a15 := uint(a12)

	const name = "bob"
	a16 := []byte(name)
	a17 := string(a16)

	a18 := []rune(name)
	a19 := string(a18)

	// 2. Custom type conversions
	a23 := MyInt(i)
	a24 := int(a23)

	// 3. Pointer operations (not conversions)
	a33 := &i
	a34 := *a33

	// 4. Unsafe conversions
	a42 := (*float64)(unsafe.Pointer(&i)) // brackets over type is needed otherwise `*float64(...)`is interpreted as de-referencing result of the float64(...) type conversion

	// 5. Channel conversions
	a53 := (<-chan int)(channel) // brackets over type is needed otherwise `<-chan` is interpreted as a a channel operation
	a54 := (chan<- int)(channel)

	// 6. Struct conversions
	a65 := Point{1, 2}
	a66 := Coord(a65)
	a67 := &Point{1, 2}
	a68 := (*Coord)(a67)

	// 7. Function conversions
	a73 := MyFunc(half)
	a74 := (MyFunc)(double)

	// 8. Interface conversions
	a86 := G(h)
	a87 := E(h) // struct to interface
	a88 := (F)(h)
	a89 := E(a86) // interface to interface
	a891 := F(a86)

	fmt.Println(a11, a12, a13, a14, a15, a16, a17, a18, a19,
		a21, a22, a23, a24,
		a31, a32, a33, a34,
		a41, a42,
		a51, a52, a53, a54,
		a61, a62, a63, a64, a65, a66, a67, a68,
		a71(i), a72(i), a73(i), a74(i),
		a81.G(), a82.E(), a83.F(), a84.E(), a85.F(), a86.G(), a87.E(), a88.F(), a89.E(), a891.F())
}
