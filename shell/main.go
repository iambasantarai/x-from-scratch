package main

import (
	"bufio"
	"fmt"
	"os"
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
		} else {
			fmt.Println(arg + ": not found")
		}
	case "echo":
		fmt.Println(strings.Join(args[1:], " "))
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
	builtins := []string{"echo", "exit", "type"}

	return slices.Contains(builtins, cmd)
}
