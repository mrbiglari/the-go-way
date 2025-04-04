package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func temporaryFilesAndDirectories() {
	check := func(e error) {
		if e != nil {
			panic(e)
		}
	}
	print := fmt.Println

	temp1, _error1 := os.CreateTemp("", "sample1")
	defer os.Remove(temp1.Name())
	check(_error1)
	print(temp1.Name())

	content := []byte("bob went to sydney")
	temp1.Write(content)

	directory, _error2 := os.MkdirTemp("", "sampledir")
	defer os.RemoveAll(directory)
	check(_error2)
	print(directory)

	temp2Path := filepath.Join(directory, "sample2")
	os.WriteFile(temp2Path, content, 0666)
}
