package main

import (
	/*	"database/sql"
		"encoding/base64"
		"encoding/json"*/
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

type check struct {
	ID      int
	HostUID int
	Command string
}

/*func exec() int {
	fmt.Println("Executing.. *NOT IMPLEMENTED YET*")

	return 0 //exit status
}*/
func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output AAH: %s\n", string(outs))
	}
}

func justDoIt() {

	cmd := exec.Command("cat", "/etc/hosts")

	var waitStatus syscall.WaitStatus
	if err := cmd.Run(); err != nil {
		printError(err)
		// Did the command fail because of an unsuccessful exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			printOutput([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
		}
	} else {
		// Command was successful
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		printOutput([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
	}
}

func main() {
	justDoIt()
}
