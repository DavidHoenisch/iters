package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: iters -c <command>")
		os.Exit(1)
	}

	commandFlag := os.Args[1]
	if commandFlag != "-c" {
		fmt.Println("Usage: iters -c <command>")
		os.Exit(1)
	}

	command := os.Args[2]
	args := os.Args[3:]

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		executeCommand(command, append(args, line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
		os.Exit(1)
	}
}

func executeCommand(command string, args []string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error executing command:", err)
	}
}
