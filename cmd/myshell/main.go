package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error reading input: %s", err)
		}

		// Handle the command
		handleCommand(command)
	}
}

func handleCommand(s string) {
	// Remove the newline character from the command
	s = strings.TrimSuffix(s, "\n")
	command := strings.Split(s, " ")[0]
	args := strings.Split(s, " ")

	switch command {
	case "echo":
		str := strings.Join(args[1:], " ")
		fmt.Fprintf(os.Stdout, "%s\n", str)
	case "exit":
		os.Exit(0)
	default:
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}
