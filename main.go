package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func main() {
	thisInvoked := path.Base(os.Args[0])
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v error: %v\n", thisInvoked, err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("usage: %v COMMAND [ARGS...]", os.Args[0])
	}
	cmd := os.Args[1]
	args := os.Args[2:]

	return platformElevate(cmd, args)
}

func plainExec(command string, args []string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Start(); err != nil {
		return err
	}

	return cmd.Wait()
}
