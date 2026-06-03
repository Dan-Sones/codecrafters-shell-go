package main

import (
	"bufio"
	"fmt"
	"os"
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

		if input == "exit" {
			os.Exit(0)
		}
		if strings.HasPrefix(input, "echo ") {
			fmt.Println(strings.TrimPrefix(input, "echo "))
		}
		if strings.HasPrefix(input, "type ") {
			postFix := strings.TrimPrefix(input, "type ")
			if slices.Contains(builtIns, postFix) {
				fmt.Printf("%s is a shell builtin \n", postFix)
			} else {
				fmt.Printf("%s: not found \n", postFix)
			}
		}
		if !slices.Contains(builtIns, strings.Split(input, " ")[0]) {
			fmt.Printf("%s: command not found \n", input)
		}
	}

}
