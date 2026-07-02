package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("✦ Welcome to the Base Converter ✦")
	fmt.Println("Type 'help' to see available commands.")

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

		switch command {
		case "convert":
			if len(parts) != 3 {
				fmt.Println("Error: usage → convert <number> <base>")
				continue
			}
			number := parts[1]
			base := strings.ToLower(parts[2])
			err := convert(number, base)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "help":
			printHelp()
		case "quit":
			fmt.Println("Goodbye! ✦")
			return
		default:
			fmt.Printf("Unknown command: %s. Type 'help' for a list of commands.\n", command)
		}
	}
}
