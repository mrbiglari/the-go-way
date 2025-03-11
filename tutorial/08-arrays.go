package main

import "fmt"

func arrays() {

	var a = [3]string{"a", "b", "c"}
	fmt.Println("emp:", a)

	var b = [...]int{100, 3: 400, 500} // if you specify the index with :, the elements in between will be zeroed.
	fmt.Println("idx:", b)

	var integers = [...]int{1, 2, 3, 4, 5} // you can also have the compiler count the number of elements for you with ...
	for i, _int := range integers {
		fmt.Println(i, _int)
	}

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
