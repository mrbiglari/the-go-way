package main

import (
	"fmt"
	"math"
)

const constString string = "Bob"

const constValueA = 0
const constValueB = 1
const constValueC = 2

const (
	constValueD = 10
	constValueE = 20
	constValueF = 30
)

const (
	constValueG = iota // 0
	constValueH        // 1
	constValueI        // 2
)

func constants() {
	fmt.Println(constString)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	fmt.Println(constValueA, constValueB, constValueC, constValueD, constValueE,
		constValueF, constValueG, constValueH, constValueI)
}
