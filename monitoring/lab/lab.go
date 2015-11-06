package main

import "os"
import "os/exec"
import "syscall"
import "fmt"
import "bytes"

func main2() {
	cmd := exec.Command("go", "version")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Print(string(cmdOutput.Bytes()))

}
func main() {
	cmd := exec.Command("cat", "8.8.1.6", "-c 1") //complete bullshit for exit code simulation
	cmdOutput := &bytes.Buffer{}
	errOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	cmd.Stderr = errOutput

	var waitStatus syscall.WaitStatus
	if err := cmd.Run(); err != nil {
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("Error: %s\n", err.Error()))
		}
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			fmt.Printf("Output: %s\n", []byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
		}
	} else {
		// Success
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		fmt.Printf("Output: %s\n", []byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
	}
	fmt.Println(string(cmdOutput.Bytes()))
	errorMsg := string(errOutput.Bytes())
	fmt.Println(waitStatus.ExitStatus())
	fmt.Println("Error: " + errorMsg)
}
