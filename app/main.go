package main

import (
	"bufio"
	"fmt"
	"os"
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

		if input == "exit" {
			os.Exit(0)
		}
		if strings.HasPrefix(input, "echo") {
			fmt.Println(input[5:])
		} else {
			fmt.Printf("%s: command not found \n", input)
		}
	}

}
