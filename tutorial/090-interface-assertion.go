package main

import (
	"fmt"
	"io"
)

// --- 1. Interface to Struct (Pointer Receiver) ---
type Printer interface {
	Print() string
}

type MyPrinter struct{}

func (p *MyPrinter) Print() string {
	return "printing"
}

var _ Printer = (*MyPrinter)(nil) // Assert that *MyPrinter implements Printer

// --- 2. Interface to Struct (Value Receiver) ---
type Stringer interface {
	String() string
}

type MyValueType struct{}

func (v MyValueType) String() string {
	return "I'm a value type"
}

var _ fmt.Stringer = MyValueType{} // Assert that MyValueType (not a pointer) implements fmt.Stringer

// --- 3. Interface to Interface (Interface Embedding) ---
type ReadWriter interface {
	io.Reader
	io.Writer
}

var _ ReadWriter = (io.ReadWriter)(nil) // Assert that io.ReadWriter implements our ReadWriter interface
