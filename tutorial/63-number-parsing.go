package main

import (
	"fmt"
	"strconv"
)

func numberParsing() {
	print := fmt.Println

	_float64, _ := strconv.ParseFloat("6.9", 64)
	print(_float64)

	_int64, _ := strconv.ParseInt("69", 0, 64) // 0 instructs to infer the base from the string
	print(_int64)

	intFromHex, _ := strconv.ParseInt("0x45", 0, 64)
	print(intFromHex)

	_uint64, _ := strconv.ParseUint("42", 0, 64)
	print(_uint64)

	number, _ := strconv.Atoi("420")
	print(number)

	_, _error := strconv.Atoi("what number?")
	print(_error)
}
