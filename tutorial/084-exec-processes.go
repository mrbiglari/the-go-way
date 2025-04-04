package main

import (
	"os"
	"os/exec"
)

func execingProcesses() {

	args := []string{"cmd", "/C", "dir"}

	cmdPath, lookError := exec.LookPath("cmd")
	if lookError != nil {
		panic(lookError)
	}
	dirCmd := exec.Command(cmdPath, args...)
	dirCmd.Env = os.Environ()
	dirCmd.Stdout = os.Stdout
	dirCmd.Stderr = os.Stderr

	if _error := dirCmd.Run(); _error != nil {
		panic(_error)
	}

	/* UNIX
		env := os.Environ()
		execErr := syscall.Exec(binary, args, env)
	    if execErr != nil {
	        panic(execErr)
	    }
	*/
}
