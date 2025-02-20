package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		user, _ := user.Current()
		hostname, _ := os.Hostname()
		cwd, _ := os.Getwd()

		PS1(user.Username, hostname, cwd)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	switch args[0] {
	case "type":
		arg := args[1]
		if isBuiltinUtil(arg) {
			fmt.Println(arg + " is a shell builtin")
		} else if path, err := exec.LookPath(arg); err == nil {
			fmt.Printf("%s is %s\n", arg, path)
		} else {
			fmt.Println(arg + ": not found")
		}
	case "echo":
		fmt.Println(strings.Join(args[1:], " "))
	case "pwd":
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory.")
		}
		fmt.Println(cwd)
	case "cd":
		path := args[1]

		if path == "~" {
			if home, err := os.UserHomeDir(); err != nil {
				fmt.Println("Error getting home directory.")
			} else {
				path = home
			}
		}

		if err := os.Chdir(path); err != nil {
			fmt.Println("cd: " + path + ": No such file or directory")
		}
	case "exit":
		os.Exit(0)
	default:
		fmt.Printf("%s: command not found\n", args[0])
	}

	return nil
}

func PS1(username, hostname, cwd string) {
	/*
		# UNIX colors
		\033[0m -> CLEAR
		\033[38;5;45;1m -> BLUE
		\033[38;5;46;1m -> GREEN
	*/
	fmt.Fprintf(os.Stdout,
		"\033[38;5;46;1m%s@%s\033[0m:\033[38;5;45;1m%s\033[0m$ ",
		username,
		hostname,
		cwd,
	)
}

func isBuiltinUtil(cmd string) bool {
	builtins := []string{"echo", "exit", "type", "pwd"}

	return slices.Contains(builtins, cmd)
}
