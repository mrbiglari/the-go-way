package main

import "fmt"

type Celsius float64

func (c Celsius) ToFahrenheit() float64 { return float64(c)*9/5 + 32 }
func (c *Celsius) Add(degrees float64)  { *c = *c + 10 } //sets original value of `Celsius`

type OurInt int

func (oi OurInt) Double() int { return int(oi) * 2 }

type SupremeInt OurInt

func (si SupremeInt) Square() int { return int((SupremeInt)(si)) * int((SupremeInt)(si)) }

type EmbeddedInt struct{ int }

func (ei EmbeddedInt) Double() int { return ei.int * 2 }

type SupremeEmbeddedInt EmbeddedInt

func (sei SupremeEmbeddedInt) Square() int { return (EmbeddedInt)(sei).int * (EmbeddedInt)(sei).int }

type WrappedInt struct{ value int }

func (wi WrappedInt) Double() int { return wi.value * 2 }

type SupremeWrappedInt WrappedInt

func (swi SupremeWrappedInt) Square() int { return (WrappedInt)(swi).value * (WrappedInt)(swi).value }

type I1 interface{ A1() string }
type I2 I1 // A2 is just a rebrand of A1 without adding any methods or semantic meaning
type I3 interface {
	I1
	A3() string
}
type S struct{}

func (a S) A1() string { return "A1" }
func (a S) A3() string { return "A3" }

func userDefinedTypes() {
	print := fmt.Println

	const temperature = 30
	var c Celsius = (Celsius)(temperature)
	print(c.ToFahrenheit())
	c.Add(10)
	print(c)

	/*
		The following explores the differences and similarities between user-defined types and related patterns in Go.
	*/

	var x int = 10 // primitive (built-in) type
	print(x + 5)

	type MyAlias = int // type alias
	var y MyAlias = 20
	print(y + 5)

	var oi OurInt = (OurInt)(30) // user-defined type
	print(int(oi)+5, oi.Double())

	var si SupremeInt = (SupremeInt)(oi) // user-defined type
	print(int(si)+5, (OurInt)(si).Double(), si.Square())

	var ei EmbeddedInt = EmbeddedInt{40} // user-defined struct type
	print(ei.int+5, ei.Double())

	var sei SupremeEmbeddedInt = (SupremeEmbeddedInt)(ei) // user-defined type
	print(sei.int+5, (EmbeddedInt)(sei).Double(), sei.Square())

	var wi WrappedInt = WrappedInt{50} // user-defined struct type
	print(wi.value+5, wi.Double())

	var swi SupremeWrappedInt = (SupremeWrappedInt)(wi) // user-defined type
	print(swi.value+5, (WrappedInt)(swi).Double(), swi.Square())

	var s S = S{}
	var i1 I1 = s
	var i2 I2 = s
	var i3 I3 = s
	print(i1.A1(), i2.A1(), i3.A1(), i3.A3())
}
