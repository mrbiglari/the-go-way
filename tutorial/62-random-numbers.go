package main

import (
	"fmt"
	"math/rand/v2"
)

func randomNumbers() {
	print := fmt.Println
	print(rand.IntN(100), rand.IntN(100))
	print(rand.Float64()*100, rand.Float64()*100)

	print(rand.Float64())                   // 0.0 <= random float   < 1.0
	print(rand.IntN(6)+10, rand.IntN(6)+10) // 6   <= random integer < 16

	seed1 := rand.NewPCG(42, 1024)
	randomGenerator1 := rand.New(seed1)
	print(randomGenerator1.IntN(100), randomGenerator1.IntN(100), randomGenerator1.IntN(100)) // 94 49 90

	seed2 := rand.NewPCG(42, 1024)
	randomGenerator2 := rand.New(seed2)
	print(randomGenerator2.IntN(100), randomGenerator2.IntN(100), randomGenerator2.IntN(100)) // 94 49 90
}
