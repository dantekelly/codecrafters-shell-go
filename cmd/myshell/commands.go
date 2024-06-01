package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Handle Commands
func handleCommand(s string) {
	// Remove the newline character from the command
	s = strings.TrimSuffix(s, "\n")
	command := strings.Split(s, " ")[0]
	args := strings.Split(s, " ")[1:]

	switch command {
	case "type":
		typeCommand(args)
	case "cd":
		cdCommand(args)
	case "pwd":
		pwdCommand()
	case "echo":
		echoCommand(args)
	case "exit":
		os.Exit(0)
	default:
		runCommand(strings.Split(s, " "))
	}
}

// Commands
func cdCommand(args []string) {
	command := args[0]

	if strings.TrimSpace(command) == "~" {
		command = os.Getenv("HOME")
	}

	if err := os.Chdir(command); err != nil {
		Print("%s: No such file or directory\n", command)
	}
}
func echoCommand(args []string) {
	str := strings.Join(args, " ")
	Print("%s\n", str)
}
func pwdCommand() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	Print("%s\n", dir)
}
func runCommand(args []string) {
	command := args[0]

	ok, p := searchEnvPath(command)
	if ok {
		executeProgram(p, args[1:])

		return
	}
	if _, err := os.Stat(command); err == nil {
		executeProgram(command, args[1:])

		return
	}

	Print("%s: command not found\n", command)
}
func typeCommand(args []string) {
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
}

// Utilities
func executeProgram(p string, args []string) {
	command := exec.Command(p, args...)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout

	err := command.Run()
	if err != nil {
		log.Fatal("[FATAL] - ", err)
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
