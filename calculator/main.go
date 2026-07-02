package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("✦ Welcome to the Go Calculator ✦")
	fmt.Println("Type 'help' to see available commands.")

	// Map of supported commands
	operations := map[string]OperationFunc{
		"add":  add,
		"sub":  sub,
		"mul":  mul,
		"div":  div,
		"help": help,
		"quit": quit,
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		command := parts[0]
		args := parts[1:]

		if op, exists := operations[command]; exists {
			err := op(args)
			if err != nil {
				fmt.Println("Error:", err)
			}
			if command == "quit" {
				break
			}
		} else {
			fmt.Printf("Unknown command: %s. Type 'help' for a list of commands.\n", command)
		}
	}
}
