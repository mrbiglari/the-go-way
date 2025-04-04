package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func directories() {
	check := func(e error) {
		if e != nil {
			panic(e)
		}
	}
	print := fmt.Println
	const tempdir = "tempdir"
	_error1 := os.Mkdir(tempdir, 0755)
	check(_error1)
	defer os.RemoveAll(tempdir)

	createEmptyFile := func(name string) {
		check(os.WriteFile(name, []byte(""), 0644))
	}

	createEmptyFile("tempdir/file1")

	_error2 := os.MkdirAll("tempdir/parent/child", 0755)
	check(_error2)

	createEmptyFile("tempdir/parent/file2")
	createEmptyFile("tempdir/parent/file3")
	createEmptyFile("tempdir/parent/child/file4")

	directories1, _error3 := os.ReadDir("tempdir/parent")
	check(_error3)
	for _, directory := range directories1 {
		print(directory.Name(), directory.IsDir())
	}

	_error4 := os.Chdir("tempdir/parent/child") // go directory two levels in
	check(_error4)
	directories2, _ := os.ReadDir(".")
	for _, directory := range directories2 {
		print(directory.Name(), directory.IsDir()) // listing directories in `tempdir/parent/child` dir
	}

	os.Chdir("../../..") // go back to where we started

	visit := func(path string, directory fs.DirEntry, _error error) error {
		if _error != nil {
			return _error
		}
		print(path, directory.Name(), directory.IsDir())
		return nil
	}
	filepath.WalkDir(tempdir, visit)
}
