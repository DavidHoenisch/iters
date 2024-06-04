package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: iters [-p] -c <command>")
		os.Exit(1)
	}

	var command string
	var args []string
	parallel := false

	if os.Args[1] == "-p" {
		if len(os.Args) < 4 || os.Args[2] != "-c" {
			fmt.Println("Usage: iters [-p] -c <command>")
			os.Exit(1)
		}
		parallel = true
		command = os.Args[3]
		args = os.Args[4:]
	} else if os.Args[1] == "-c" {
		command = os.Args[2]
		args = os.Args[3:]
	} else {
		fmt.Println("Usage: iters [-p] -c <command>")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	if parallel {
		var wg sync.WaitGroup
		for scanner.Scan() {
			line := scanner.Text()
			wg.Add(1)
			go func(line string) {
				defer wg.Done()
				executeCommand(command, args, line)
			}(line)
		}
		wg.Wait()
	} else {
		for scanner.Scan() {
			line := scanner.Text()
			executeCommand(command, args, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
		os.Exit(1)
	}
}

func executeCommand(command string, args []string, input string) {
	var cmd *exec.Cmd

	if strings.Contains(command, "{}") {
		// Replace the placeholder with the input
		command = strings.ReplaceAll(command, "{}", input)
		// Split the command into the base command and its arguments
		parts := strings.Fields(command)
		cmd = exec.Command(parts[0], parts[1:]...)
	} else {
		// Append the input to the end of the arguments
		cmd = exec.Command(command, append(args, input)...)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error executing command:", err)
	}
}
