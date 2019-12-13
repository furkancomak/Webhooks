package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

func Execute(dirPath string, name string, arg ...string) (string, int) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	os.Stderr = w
	cmd := exec.Command(name, arg...)
	if dirPath != "" {
		cmd.Dir = dirPath
	}
	stdin, err := cmd.StdinPipe()

	if err != nil {
		fmt.Println(err)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	var waitStatus syscall.WaitStatus
	var exitCode = 0
	if err = cmd.Start(); err != nil {
		fmt.Println("An error occured: ", err)
		// Did the command fail because of an unsuccessful exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			exitCode = waitStatus.ExitStatus()
		}
	}

	stdin.Close()
	cmd.Wait()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	return string(out), exitCode
}
