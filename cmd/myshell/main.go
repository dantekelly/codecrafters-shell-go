package main

import (
	"bufio"
	"fmt"
	"os"
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

func handleCommand(command string) {
	// Remove the newline character from the command
	command = command[:len(command)-1]

	switch command {
	default:
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}
