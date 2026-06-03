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

	for {
		fmt.Print("$ ")
		input, err := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		supportedCommands := []string{}

		if input == "exit" {
			os.Exit(0)
		}
		if !slices.Contains(supportedCommands, input) {
			fmt.Printf("%s: command not found \n", input)
		}
	}

}
