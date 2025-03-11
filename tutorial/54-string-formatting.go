package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func printfn(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func stringFormatting() {

	dot := point{x: 10, y: 100}

	printfn("struct: %v", dot)  // struct: {10 100}
	printfn("struct: %+v", dot) // struct: {x:10 y:100}
	printfn("struct: %#v", dot) // struct: main.point{x:10, y:100}

	printfn("type: %T", dot) // type: main.point

	printfn("bool: %t", true) // bool: true

	printfn("int: %d", 123) // int: 123

	printfn("binary: %b", 4) // binary: 100

	printfn("character: %c", 33) // character: !

	printfn("hex: %x", 456) // hex: 1c8

	printfn("float: %f", 12.453)   // float: 12.453000
	printfn("float: %.2f", 12.453) // float: 12.45
	printfn("float: %e", 12.45)    // float: 1.245000e+01
	printfn("float: %E", 12.45)    // float: 1.245000E+01

	printfn("string: %s", "Bob") // string: Bob
	printfn("string: %q", "Bob") // string: "Bob"
	printfn("string: %x", "Bob") // string: 426f62

	printfn("pointer: %p", &dot) // pointer: 0xc00000a080 (address will vary)

	printfn("width: |%6d|%6d|", 12, 345)         // width1: |    12|   345|
	printfn("width: |%6.2f|%6.2f|", 1.2, 3.45)   // width2: |  1.20|  3.45|
	printfn("width: |%-6.2f|%-6.2f|", 1.2, 3.45) // width3: |1.20  |3.45  |
	printfn("width: |%6s|%6s|", "foo", "b")      // width4: |   foo|     b|
	printfn("width: |%-6s|%-6s|", "foo", "b")    // width5: |foo   |b     |

	text := fmt.Sprintf("sprintf: a %s\n", "book")
	fmt.Println(text) // sprintf: a book

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error") // io: an error
}
