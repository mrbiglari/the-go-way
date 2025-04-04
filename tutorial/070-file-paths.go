package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func filePaths() {
	print := fmt.Println

	path := filepath.Join("dir1", "dir2", "filename")
	print(path)

	print(filepath.Join("dir1//", "filename"))
	print(filepath.Join("dir1/../dir2", "filename")) //  Join() normalizes paths by removing superfluous separators and directory changes

	print(filepath.Dir(path))
	print(filepath.Base(path))

	print(filepath.IsAbs("dir/filename"))
	print(filepath.IsAbs("/dir/filename"))

	filename := "bob.json"

	extension := filepath.Ext(filename)
	print(extension)
	print(strings.TrimSuffix(filename, extension))

	relativePath1, _error1 := filepath.Rel("a/b", "a/b/t/file")
	if _error1 != nil {
		panic(_error1)
	}
	print(relativePath1)

	relativePath2, _ := filepath.Rel("a/b", "a/c/t/file")
	print(relativePath2)
}
