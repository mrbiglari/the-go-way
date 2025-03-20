package main

import (
	"fmt"
	"os/exec"
)

func spawningProcesses() {
	check := func(_error error) {
		if _error != nil {
			panic(_error)
		}
	}
	print := fmt.Println

	dateCmd := exec.Command("cmd", "/C", "date", "/T") // Unix equivalent:  exec.Command("date")
	dateInBytes, _error1 := dateCmd.Output()           // runs and returns its standard output
	check(_error1)
	print(string(dateInBytes))

	_, _error := exec.Command("xyz").Output()
	if _error != nil {
		switch errorType := _error.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", _error)
		case *exec.ExitError:
			fmt.Println("command exit rc=:", errorType.ExitCode())
		default:
			panic(_error)
		}
	}

	findstrCmd := exec.Command("cmd", "/C", "findstr", "hello", "01-hello-world.go")
	output, _error2 := findstrCmd.Output()
	check(_error2)
	print("match:", string(output))

	/* UNIX
		grepCmd := exec.Command("grep", "hello")

		grepIn, _ := grepCmd.StdinPipe()
		grepOut, _ := grepCmd.StdoutPipe()
		grepCmd.Start()
		grepIn.Write([]byte("hello grep\ngoodbye grep"))
		grepIn.Close()
		grepBytes, _ := io.ReadAll(grepOut)
		grepCmd.Wait()

		fmt.Println("> grep hello")
		fmt.Println(string(grepBytes))

		lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	    lsOut, err := lsCmd.Output()
	    if err != nil {
	        panic(err)
	    }
	    fmt.Println("> ls -a -l -h")
	    fmt.Println(string(lsOut))
	*/
}
