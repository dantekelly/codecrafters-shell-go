package main

import (
	"bufio"
	"os"
)

func main() {
	handleInput()
}

func handleInput() {
	reader := bufio.NewReader(os.Stdin)
	for {
		Print("$ ")

		// Wait for user input
		command, err := reader.ReadString('\n')
		if err != nil {
			Print("Error reading input: %s", err)
		}

		// Handle the command
		handleCommand(command)
	}
}
