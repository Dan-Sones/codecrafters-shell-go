package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	reader := bufio.NewReader(os.Stdin)
	builtIns := []string{"exit", "echo", "type"}

	for {
		fmt.Print("$ ")
		input, err := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		parts := strings.Split(input, " ")
		command := parts[0]
		args := parts[1:]

		switch command {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Println(strings.Join(args, " "))
		case "type":
			if slices.Contains(builtIns, args[0]) {
				fmt.Printf("%s is a shell builtin \n", args[0])
			} else if path, _ := exec.LookPath(args[0]); path != "" {
				fmt.Printf("%s is %s\n", args[0], path)
			} else {
				fmt.Printf("%s: not found\n", args[0])
			}
		default:
			handleAmbiguousArgs(command, args)

		}
	}
}

func handleAmbiguousArgs(command string, args []string) {
	// first check to see if the program is on the path
	if path, _ := exec.LookPath(command); path != "" {
		cmd := exec.Command(command, args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("%s: command not found \n", command)
	}
}
