package main

import (
	"bufio"
	"fmt"
	"os"
)

func writingFiles() {
	check := func(_error error) {
		if _error != nil {
			panic(_error)
		}
	}
	print := fmt.Println

	const contentString = "package main\n"
	content := []byte(contentString)

	const tempFile1 = "temp1.txt"
	_error := os.WriteFile(tempFile1, content, 0644)
	check(_error)

	const tempFile2 = "temp2.txt"
	file, _ := os.Create(tempFile2)
	defer file.Close()

	bytesWrittenCount1, _ := file.Write(content)
	print(bytesWrittenCount1)

	bytesWrittenCount2, _ := file.WriteString(contentString)
	print(bytesWrittenCount2)

	file.Sync() // force data out of buffers

	writer := bufio.NewWriter(file)
	bytesWrittenCount3, _ := writer.WriteString(contentString)
	print(bytesWrittenCount3)

	writer.Flush() // force data out of buffers
}
