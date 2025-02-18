package main

import "fmt"

func _for() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 8; j < 10; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for z := range 5 {
		fmt.Println(z)
	}

	for t := range 6 {
		if t%2 == 0 {
			continue
		}
		fmt.Println(t)
	}

	iterables := []string{"a", "b", "c", "d", "e"}
	for index := range iterables {
		fmt.Println(iterables[index])
	}

	for index, value := range iterables {
		fmt.Println(index, value)
	}
}
