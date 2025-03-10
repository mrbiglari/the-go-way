package main

import (
	"cmp"
	"fmt"
	"slices"
)

func sortingByFunctions() {
	fruits := []string{"banana", "peach", "kiwi"}
	compareFunc := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	fmt.Println("Fruits before sorting:", fruits)
	slices.SortFunc(fruits, compareFunc)
	fmt.Println("Fruits after sorting:", fruits)

	type person struct {
		name string
		age  int
	}
	people := []person{{name: "Bob", age: 21}, {name: "Sydney", age: 35}, {name: "Amy", age: 24}}

	fmt.Println("People before sorting:", people)
	slices.SortFunc(people, func(a, b person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println("People after sorting:", people)
}
