package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

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

func searchEnvPath(command string) (bool, string) {
	paths := strings.Split(os.Getenv("PATH"), ":")

	for _, path := range paths {
		fp := filepath.Join(path, command)
		if _, err := os.Stat(fp); err == nil {
			return true, fp
		}
	}

	return false, ""
}

func handleCommand(s string) {
	// Remove the newline character from the command
	s = strings.TrimSuffix(s, "\n")
	command := strings.Split(s, " ")[0]
	args := strings.Split(s, " ")[1:]

	switch command {
	case "type":
		if len(args) >= 1 {
			switch args[0] {
			case "type", "echo", "exit":
				Print("%s is a shell builtin\n", args[0])
			default:
				ok, p := searchEnvPath(args[0])
				if ok {
					Print("%s is %s\n", args[0], p)
					return
				}

				Print("%s not found\n", args[0])
			}

			return
		}

		Print("type: missing argument\n")
	case "echo":
		str := strings.Join(args, " ")
		Print("%s\n", str)
	case "exit":
		os.Exit(0)
	default:
		Print("%s: command not found\n", command)
	}
}
