package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readingFiles() {
	check := func(_error error) {
		if _error != nil {
			panic(_error)
		}
	}
	print := fmt.Println

	content, _error := os.ReadFile("01-hello-world.go")
	check(_error)
	print(string(content))

	file, _ := os.Open("01-hello-world.go") // first line of file: `package main`
	defer file.Close()

	const packageWordLength = 7
	chunk1 := make([]byte, packageWordLength)
	length1, _ := file.Read(chunk1)
	print(string(chunk1[:length1])) // package

	offset, _ := file.Seek(packageWordLength+1, io.SeekStart)
	print(offset)

	const mainWordLength = 5
	chunk2 := make([]byte, mainWordLength)
	length2, _ := file.Read(chunk2)
	print(string(chunk2[:length2])) // main

	file.Seek(2, io.SeekCurrent)
	file.Seek(-4, io.SeekEnd)

	const ageWordLength = 3
	file.Seek(4 /* pack'age' */, io.SeekStart)
	chunk3 := make([]byte, ageWordLength)
	bytesCopiedCount, _ := io.ReadAtLeast(file, chunk3, ageWordLength)
	print(bytesCopiedCount)
	print(string(chunk3)) // age

	file.Seek(0, io.SeekStart) // no built-in rewind but this serves as a reset

	reader := bufio.NewReader(file)
	peek, _ := reader.Peek(packageWordLength)
	print(string(peek)) // package
}
