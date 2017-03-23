// Found here: https://nathanleclaire.com/blog/2014/12/29/shelled-out-commands-in-golang/ by Nathan Leclaire
// Modified by Alexander Sosna

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmdName := "echo"
	cmdArgs := []string{"hello", "world"}

	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("docker build out | %s\n", scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		os.Exit(1)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		os.Exit(1)
	}
}
